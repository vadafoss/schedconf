package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	config "k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/kubernetes/pkg/scheduler/apis/config/scheme"
	"k8s.io/kubernetes/pkg/scheduler/apis/config/validation"
	// "k8s.io/kubernetes/cmd/kube-scheduler/app/options"
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
