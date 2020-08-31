package snapshot

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/pkg/k8sutil"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/utils"
	"github.com/stackrox/rox/sensor/upgrader/common"
	"github.com/stackrox/rox/sensor/upgrader/upgradectx"
	v1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
)

var (
	log = logging.LoggerForModule()

	jsonSeparator = []byte("\x00")
)

type snapshotter struct {
	ctx  *upgradectx.UpgradeContext
	opts Options
}

func (s *snapshotter) SnapshotState() ([]k8sutil.Object, error) {
	coreV1Client := s.ctx.ClientSet().CoreV1()

	snapshotSecret, err := coreV1Client.Secrets(common.Namespace).Get(s.ctx.Context(), secretName, metav1.GetOptions{})
	if k8sErrors.IsNotFound(err) {
		snapshotSecret = nil
		err = nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "retrieving state snapshot secret")
	}

	// Ignore the snapshot secret if it doesn't belong to this upgrade process. This means storing will fail if
	// requested, which is okay.
	if snapshotSecret != nil && s.ctx.IsProcessStateObject(snapshotSecret) {
		log.Info("Matching state snapshot secret found, not creating a new one")
		return s.stateFromSecret(snapshotSecret)
	}

	if s.opts.MustExist {
		return nil, errors.New("state snapshot secret does not exist")
	}

	objects, snapshotSecret, err := s.createStateSnapshot()
	if err != nil {
		return nil, errors.Wrap(err, "snapshotting state")
	}

	if s.opts.Store {
		_, err = coreV1Client.Secrets(common.Namespace).Create(s.ctx.Context(), snapshotSecret, metav1.CreateOptions{})
		if err != nil {
			return nil, errors.Wrap(err, "creating state snapshot secret")
		}
	}
	return objects, nil
}

func (s *snapshotter) stateFromSecret(secret *v1.Secret) ([]k8sutil.Object, error) {
	if processID := secret.Labels[common.UpgradeProcessIDLabelKey]; processID != s.ctx.ProcessID() {
		return nil, errors.Errorf("state snapshot secret belongs to wrong upgrade process %q, expected %s", processID, s.ctx.ProcessID())
	}

	gzData := secret.Data[secretDataName]
	if len(gzData) == 0 {
		return nil, errors.New("state snapshot secret contains no relevant data")
	}

	gzReader, err := gzip.NewReader(bytes.NewReader(gzData))
	if err != nil {
		return nil, errors.Wrap(err, "creating gzip readere for state snapshot data")
	}

	allObjBytes, err := ioutil.ReadAll(gzReader)
	if err != nil {
		return nil, errors.Wrap(err, "reading compressed state snapshot data")
	}
	if err := gzReader.Close(); err != nil {
		return nil, errors.Wrap(err, "reading compressed state snapshot data")
	}

	if len(allObjBytes) == 0 {
		return nil, nil
	}

	objBytes := bytes.Split(allObjBytes, jsonSeparator)

	universalDeserializer := s.ctx.UniversalDecoder()

	result := make([]k8sutil.Object, 0, len(objBytes))
	for _, serialized := range objBytes {
		runtimeObj, _, err := universalDeserializer.Decode(serialized, nil, nil)
		if err != nil {
			return nil, errors.Wrap(err, "could not deserialize object in stored snapshot")
		}
		obj, _ := runtimeObj.(k8sutil.Object)
		if obj == nil {
			return nil, errors.Errorf("object of kind %v does not have object metadata", runtimeObj.GetObjectKind().GroupVersionKind())
		}
		result = append(result, obj)
	}

	return result, nil
}

func (s *snapshotter) createStateSnapshot() ([]k8sutil.Object, *v1.Secret, error) {
	objs, err := s.ctx.ListCurrentObjects()
	if err != nil {
		return nil, nil, err
	}

	byteSlices := make([][]byte, 0, len(objs))
	jsonSerializer := json.NewSerializer(json.DefaultMetaFactory, nil, nil, false)
	for _, obj := range objs {
		var buf bytes.Buffer
		if err := jsonSerializer.Encode(obj, &buf); err != nil {
			return nil, nil, errors.Wrapf(err, "marshaling object of kind %v to JSON", obj.GetObjectKind().GroupVersionKind())
		}
		byteSlices = append(byteSlices, buf.Bytes())
	}

	var compressedData bytes.Buffer
	gzipWriter, err := gzip.NewWriterLevel(&compressedData, gzip.BestCompression)
	if err != nil {
		return nil, nil, utils.Should(err) // level is valid, so expect no error
	}
	if _, err := gzipWriter.Write(bytes.Join(byteSlices, jsonSeparator)); err != nil {
		return nil, nil, utils.Should(err)
	}
	if err := gzipWriter.Close(); err != nil {
		return nil, nil, utils.Should(err)
	}

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: common.Namespace,
			Name:      secretName,
		},
		Type: v1.SecretTypeOpaque,
		Data: map[string][]byte{
			secretDataName: compressedData.Bytes(),
		},
	}
	s.ctx.AnnotateProcessStateObject(secret)

	return objs, secret, nil
}
