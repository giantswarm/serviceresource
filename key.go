package service

import (
	"github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/microerror"
)

const (
	serviceNameMaster = "master"
	serviceNameWorker = "worker"
)

func clusterID(customObject interface{}) (string, error) {
	// try aws
	awsConfig , err := toAWSConfig(customObject)
	if err == nil {
		return awsConfig.Spec.Cluster.ID, nil
	} else if !IsWrongTypeError(err) {
		return "", microerror.Maskf(err, "unexpected error when converting custom object")
	}

	// try azure
	azureConfig , err := toAzureConfig(customObject)
	if err == nil {
		return azureConfig.Spec.Cluster.ID, nil
	} else if !IsWrongTypeError(err) {
		return "", microerror.Maskf(err, "unexpected error when converting custom object")
	}

	// try kvm
	kvmConfig , err := toKVMConfig(customObject)
	if err == nil {
		return kvmConfig.Spec.Cluster.ID, nil
	} else if !IsWrongTypeError(err) {
		return "", microerror.Maskf(err, "unexpected error when converting custom object")
	}

	return "", microerror.Maskf(err, "unknown custom object type")
}

func clusterNamespace(clusterID string) string{
	return clusterID
}

func isDeleted(customObject v1alpha1.KVMConfig) bool {
	return customObject.GetDeletionTimestamp() != nil
}
func isKVM(customObject interface{}) bool  {
	_ , err := toKVMConfig(customObject)
	if err == nil {
		return true
	} else  {
		return false
	}
}

func toAWSConfig (v interface{}) (v1alpha1.AWSConfig, error) {
	if v == nil {
		return v1alpha1.AWSConfig{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.AWSConfig{}, v)
	}

	customObjectPointer, ok := v.(*v1alpha1.AWSConfig)
	if !ok {
		return v1alpha1.AWSConfig{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.AWSConfig{}, v)
	}
	customObject := *customObjectPointer

	customObject = *customObject.DeepCopy()

	return customObject, nil
}

func toKVMConfig(v interface{}) (v1alpha1.KVMConfig, error)  {
	if v == nil {
		return v1alpha1.KVMConfig{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.KVMConfig{}, v)
	}

	customObjectPointer, ok := v.(*v1alpha1.KVMConfig)
	if !ok {
		return v1alpha1.KVMConfig{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.KVMConfig{}, v)
	}
	customObject := *customObjectPointer

	customObject = *customObject.DeepCopy()

	return customObject, nil

}

func toAzureConfig(v interface{}) (v1alpha1.AzureConfig, error) {
	if v == nil {
		return v1alpha1.AzureConfig{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.AzureConfig{}, v)
	}

	customObjectPointer, ok := v.(*v1alpha1.AzureConfig)
	if !ok {
		return v1alpha1.AzureConfig{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.AzureConfig{}, v)
	}
	customObject := *customObjectPointer

	customObject = *customObject.DeepCopy()

	return customObject, nil
}