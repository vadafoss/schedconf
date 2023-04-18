package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v3"

	scheduler_config "k8s.io/kubernetes/pkg/scheduler/apis/config/latest"
)

func main() {
	config, err := scheduler_config.Default()
	if err != nil {
		fmt.Printf("couldn't create scheduler config: %v", err)
		panic(1)
	}

	b, err := yaml.Marshal(config)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
