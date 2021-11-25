package augmentedobjs

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/booleanpolicy/evaluator/pathutil"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/signature"
	"github.com/stackrox/rox/pkg/utils"
)

const (
	// CompositeFieldCharSep is the separating character used when we create a composite field.
	CompositeFieldCharSep = "\t"
)

var (
	log = logging.LoggerForModule()
)

// TODO(dhaus): Declare a interface / function alias here to verify the image signature (like what is the return value)
type verifySignatureOfDeploymentImages = func(deployment *storage.Deployment, keysToVerify []string) (string, bool)

func findMatchingContainerIdxForProcess(deployment *storage.Deployment, process *storage.ProcessIndicator) (int, error) {
	for i, container := range deployment.GetContainers() {
		if container.GetName() == process.GetContainerName() {
			return i, nil
		}
	}
	return 0, errors.Errorf("indicator %s could not be matched (container name %s not found in deployment %s/%s/%s",
		process.GetSignal().GetExecFilePath(), process.GetContainerName(), deployment.GetClusterId(), deployment.GetNamespace(), deployment.GetName())

}

// ConstructDeploymentWithProcess constructs an augmented deployment with process information.
func ConstructDeploymentWithProcess(deployment *storage.Deployment,
	images []*storage.Image,
	verifierFactory signature.VerifierFactory,
	process *storage.ProcessIndicator,
	processNotInBaseline bool) (*pathutil.AugmentedObj, error) {
	obj, err := ConstructDeployment(deployment, images, verifierFactory)
	if err != nil {
		return nil, err
	}
	augmentedProcess, err := ConstructProcess(process, processNotInBaseline)
	if err != nil {
		return nil, err
	}

	matchingContainerIdx, err := findMatchingContainerIdxForProcess(deployment, process)
	if err != nil {
		return nil, err
	}
	err = obj.AddAugmentedObjAt(
		augmentedProcess,
		pathutil.FieldStep("Containers"), pathutil.IndexStep(matchingContainerIdx), pathutil.FieldStep(processAugmentKey),
	)
	if err != nil {
		return nil, utils.Should(err)
	}
	return obj, nil
}

// ConstructKubeResourceWithEvent constructs an augmented deployment with kube event information.
func ConstructKubeResourceWithEvent(kubeResource interface{}, event *storage.KubernetesEvent) (*pathutil.AugmentedObj, error) {
	obj, isAugmented := kubeResource.(*pathutil.AugmentedObj)
	if !isAugmented {
		if !supportedKubeResourceForEvent(kubeResource) {
			return nil, errors.Errorf("unsupported kubernetes resource type %T for event detection", kubeResource)
		}
		obj = pathutil.NewAugmentedObj(kubeResource)
	}

	if err := obj.AddPlainObjAt(event, pathutil.FieldStep(kubeEventAugKey)); err != nil {
		return nil, utils.Should(err)
	}
	return obj, nil
}

func supportedKubeResourceForEvent(obj interface{}) bool {
	switch obj.(type) {
	case *storage.Deployment:
		return true
	default:
		return false
	}
}

// ConstructProcess constructs an augmented process.
func ConstructProcess(process *storage.ProcessIndicator, processNotInBaseline bool) (*pathutil.AugmentedObj, error) {
	augmentedProcess := pathutil.NewAugmentedObj(process)
	err := augmentedProcess.AddPlainObjAt(
		&baselineResult{NotInBaseline: processNotInBaseline},
		pathutil.FieldStep(baselineResultAugmentKey),
	)
	if err != nil {
		return nil, errors.Wrap(err, "adding process baseline result to process")
	}
	return augmentedProcess, nil
}

// ConstructKubeEvent constructs an augmented kubernetes event.
func ConstructKubeEvent(event *storage.KubernetesEvent) *pathutil.AugmentedObj {
	return pathutil.NewAugmentedObj(event)
}

// ConstructAuditEvent constructs an augmented kubernetes event.
func ConstructAuditEvent(event *storage.KubernetesEvent, isImpersonated bool) (*pathutil.AugmentedObj, error) {
	augmentedProcess := pathutil.NewAugmentedObj(event)

	err := augmentedProcess.AddPlainObjAt(
		&impersonatedEventResult{IsImpersonatedUser: isImpersonated},
		pathutil.FieldStep(impersonatedEventResultKey),
	)
	if err != nil {
		return nil, errors.Wrap(err, "adding is impersonated user result to audit event")
	}
	return augmentedProcess, nil
}

// ConstructNetworkFlow constructs an augmented network flow.
func ConstructNetworkFlow(flow *NetworkFlowDetails) (*pathutil.AugmentedObj, error) {
	augmentedFlow := pathutil.NewAugmentedObj(flow)
	return augmentedFlow, nil
}

// ConstructDeploymentWithNetworkFlowInfo constructs an augmented object with deployment and network flow.
func ConstructDeploymentWithNetworkFlowInfo(
	deployment *storage.Deployment,
	images []*storage.Image,
	verifierFactory signature.VerifierFactory,
	flow *NetworkFlowDetails,
) (*pathutil.AugmentedObj, error) {
	obj, err := ConstructDeployment(deployment, images, verifierFactory)
	if err != nil {
		return nil, err
	}
	augmentedFlow, err := ConstructNetworkFlow(flow)
	if err != nil {
		return nil, err
	}

	err = obj.AddAugmentedObjAt(augmentedFlow, pathutil.FieldStep(networkFlowAugKey))
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// ConstructDeployment constructs the augmented deployment object.
func ConstructDeployment(deployment *storage.Deployment, images []*storage.Image, verifierFactory signature.VerifierFactory) (*pathutil.AugmentedObj, error) {
	obj := pathutil.NewAugmentedObj(deployment)
	if len(images) != len(deployment.GetContainers()) {
		return nil, errors.Errorf("deployment %s/%s had %d containers, but got %d images",
			deployment.GetNamespace(), deployment.GetName(), len(deployment.GetContainers()), len(images))
	}
	for i, image := range images {
		augmentedImg, err := ConstructImage(image)
		if err != nil {
			return nil, err
		}
		err = obj.AddAugmentedObjAt(
			augmentedImg,
			pathutil.FieldStep("Containers"), pathutil.IndexStep(i), pathutil.FieldStep(imageAugmentKey),
		)
		if err != nil {
			return nil, utils.Should(err)
		}
	}

	for idx, container := range deployment.GetContainers() {
		for i, env := range container.GetConfig().GetEnv() {
			envVarObj := &envVar{EnvVar: fmt.Sprintf("%s%s%s%s%s", env.GetEnvVarSource(), CompositeFieldCharSep, env.GetKey(), CompositeFieldCharSep, env.GetValue())}
			err := obj.AddPlainObjAt(
				envVarObj,
				pathutil.FieldStep("Containers"), pathutil.IndexStep(idx), pathutil.FieldStep("Config"),
				pathutil.FieldStep("Env"), pathutil.IndexStep(i), pathutil.FieldStep(envVarAugmentKey),
			)

			if err != nil {
				return nil, utils.Should(err)
			}
		}
	}

	verifiedImageSignatureResult := &verifiedImageSignatureKeyResult{VerifiedImageSignatureKeys: map[string]string{}}

	if verifierFactory != nil {
		log.Infof("Factory is set during creation, will create and verify the signature")
		dv, err := verifierFactory.DeploymentVerifier(deployment)
		if err != nil {
			log.Errorf("Found error during creation of delpoyment verifier: %v", err)
			return nil, err
		}
		verifiedKeys, verified, err := dv.VerifyImages()
		if err != nil {
			log.Errorf("Found error during verification of images: %v", err)
			return nil, err
		}
		if verified {
			log.Infof("The image is verified and the verified keys are %v", verifiedKeys)
			verifiedImageSignatureResult.VerifiedImageSignatureKeys = verifiedKeys
		}
	}

	if err := obj.AddPlainObjAt(verifiedImageSignatureResult,
		pathutil.FieldStep(imageSignatureAugmentKey)); err != nil {
		return nil, err
	}

	return obj, nil
}

// ConstructImage constructs the augmented image object.
func ConstructImage(image *storage.Image) (*pathutil.AugmentedObj, error) {
	obj := pathutil.NewAugmentedObj(image)

	// Since policies query for Dockerfile Line as a single compound field, we simulate it by creating a "composite"
	// dockerfile line under each layer.
	for i, layer := range image.GetMetadata().GetV1().GetLayers() {
		lineObj := &dockerfileLine{Line: fmt.Sprintf("%s%s%s", layer.GetInstruction(), CompositeFieldCharSep, layer.GetValue())}
		err := obj.AddPlainObjAt(
			lineObj,
			pathutil.FieldStep("Metadata"), pathutil.FieldStep("V1"), pathutil.FieldStep("Layers"),
			pathutil.IndexStep(i), pathutil.FieldStep(dockerfileLineAugmentKey),
		)
		if err != nil {
			return nil, utils.Should(err)
		}
	}

	// Since policies query for component and version as a single compound field, we simulate it by creating a
	// "composite" component and version field.
	for i, component := range image.GetScan().GetComponents() {
		compAndVersionObj := &componentAndVersion{
			ComponentAndVersion: fmt.Sprintf("%s%s%s", component.GetName(), CompositeFieldCharSep, component.GetVersion()),
		}
		err := obj.AddPlainObjAt(
			compAndVersionObj,
			pathutil.FieldStep("Scan"), pathutil.FieldStep("Components"), pathutil.IndexStep(i),
			pathutil.FieldStep(componentAndVersionAugmentKey),
		)
		if err != nil {
			return nil, utils.Should(err)
		}
	}
	return obj, nil
}
