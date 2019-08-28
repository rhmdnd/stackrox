package bundle

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/pkg/utils"
	"github.com/stackrox/rox/sensor/upgrader/common"
	"github.com/stackrox/rox/sensor/upgrader/k8sobjects"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func readFile(openFn OpenFunc) ([]byte, error) {
	reader, err := openFn()
	if err != nil {
		return nil, err
	}
	defer utils.IgnoreError(reader.Close)

	return ioutil.ReadAll(reader)
}

func createDynamicObject(objDesc common.DynamicBundleObjectDesc, bundleContents Contents) (k8sobjects.Object, error) {
	dataMap := make(map[string][]byte)

	// If the object is optional, perform a first pass to check if all files are present (and return nil if files are
	// missing).
	if objDesc.Optional {
		for _, fileName := range objDesc.Files {
			if bundleContents.File(fileName) == nil {
				return nil, nil
			}
		}
	}

	for _, fileName := range objDesc.Files {
		openFn := bundleContents.File(fileName)
		if openFn == nil {
			// optional case already handled above
			return nil, errors.Errorf("required file %s not found in bundle", fileName)
		}

		fileData, err := readFile(openFn)
		if err != nil {
			return nil, errors.Wrapf(err, "reading file %s from bundle", fileName)
		}
		dataMap[fileName] = fileData
	}

	var obj k8sobjects.Object

	switch objDesc.Kind {
	case common.OpaqueSecret:
		obj = &v1.Secret{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "v1",
				Kind:       "Secret",
			},
			Type: v1.SecretTypeOpaque,
			Data: dataMap,
		}
	case common.ConfigMap:
		{
			strData := make(map[string]string)
			for k, v := range dataMap {
				if utf8.Valid(v) {
					strData[k] = string(v)
					delete(dataMap, k)
				}
			}

			obj = &v1.ConfigMap{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "v1",
					Kind:       "ConfigMap",
				},
				Data:       strData,
				BinaryData: dataMap,
			}
		}
	default:
		return nil, errors.Errorf("unknown dynamic bundle object kind %v", objDesc.Kind)
	}

	if obj.GetLabels() == nil {
		obj.SetLabels(make(map[string]string))
	}
	obj.GetLabels()[common.UpgradeResourceLabelKey] = common.UpgradeResourceLabelValue
	return obj, nil
}

func createDynamicObjects(bundleContents Contents) ([]k8sobjects.Object, error) {
	var allObjects []k8sobjects.Object
	for _, objDesc := range common.DynamicBundleObjects {
		obj, err := createDynamicObject(objDesc, bundleContents)
		if err != nil {
			return nil, errors.Wrapf(err, "could not instantiate dynamic bundle object %s", objDesc.Name)
		}
		if obj != nil {
			allObjects = append(allObjects, obj)
		}
	}
	return allObjects, nil
}
