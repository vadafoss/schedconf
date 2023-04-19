package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"

	config "k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/kubernetes/pkg/scheduler/apis/config/scheme"
	"k8s.io/kubernetes/pkg/scheduler/apis/config/validation"
)

func main() {
	// cfg, err := scheduler_config.Default()
	// if err != nil {
	// 	fmt.Printf("couldn't create scheduler config: %v", err)
	// 	panic(1)
	// }

	file := "sample-config.yaml"
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

	// cfg.Profiles = cfgObj.Profiles

	if err := validation.ValidateKubeSchedulerConfiguration(cfgObj); err != nil {
		panic(err)
	}

	// b1, err := yaml.Marshal(cfg)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(b1))

	b2, err := yaml.Marshal(cfgObj)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b2))
}
