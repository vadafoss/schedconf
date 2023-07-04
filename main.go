package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"
	config "k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/kubernetes/pkg/scheduler/apis/config/scheme"
	"k8s.io/kubernetes/pkg/scheduler/apis/config/validation"

	// "k8s.io/kubernetes/cmd/kube-scheduler/app/options"
	configv1beta3 "k8s.io/kubernetes/pkg/scheduler/apis/config/v1beta3"
)

func main() {

	// Uncomment this line to get the default config
	// schedConfig, err := scheduler_config.Default()
	// if err != nil {
	// 	panic(err)
	// }

	// Uncomment and use this function to write the config file into correct yaml format
	// err = options.LogOrWriteConfig("test-sched-config.yaml", schedConfig, schedConfig.Profiles)
	// if err != nil {
	// 	panic(err)
	// }

	file := "test-config-plugin-disabled.yaml"
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	obj, gvk, err := scheme.Codecs.UniversalDecoder().Decode(data, nil, nil)
	if err != nil {
		panic(err)
	}
	cfgObj, ok := obj.(*config.KubeSchedulerConfiguration)
	if !ok {
		panic(fmt.Errorf("couldn't decode as KubeSchedulerConfiguration, got %s: ", gvk))
	}

	cfgObj.TypeMeta.APIVersion = gvk.GroupVersion().String()

	if err := validation.ValidateKubeSchedulerConfiguration(cfgObj); err != nil {
		panic(err)
	}

	data, err = yaml.Marshal(cfgObj)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}

func LoadConfigFromFile(logger klog.Logger, file string) (*config.KubeSchedulerConfiguration, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return loadConfig(logger, data)
}

func loadConfig(logger klog.Logger, data []byte) (*config.KubeSchedulerConfiguration, error) {
	// The UniversalDecoder runs defaulting and returns the internal type by default.
	obj, gvk, err := scheme.Codecs.UniversalDecoder().Decode(data, nil, nil)
	if err != nil {
		return nil, err
	}
	if cfgObj, ok := obj.(*config.KubeSchedulerConfiguration); ok {
		// We don't set this field in pkg/scheduler/apis/config/{version}/conversion.go
		// because the field will be cleared later by API machinery during
		// conversion. See KubeSchedulerConfiguration internal type definition for
		// more details.
		cfgObj.TypeMeta.APIVersion = gvk.GroupVersion().String()
		switch cfgObj.TypeMeta.APIVersion {
		case configv1beta3.SchemeGroupVersion.String():
			logger.Info("KubeSchedulerConfiguration v1beta3 is deprecated in v1.26, will be removed in v1.29")
		}
		return cfgObj, nil
	}
	return nil, fmt.Errorf("couldn't decode as KubeSchedulerConfiguration, got %s: ", gvk)
}
