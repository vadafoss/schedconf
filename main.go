package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"

	config "k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/kubernetes/pkg/scheduler/apis/config/scheme"
)

func main() {
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

	b2, err := yaml.Marshal(cfgObj)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b2))
}
