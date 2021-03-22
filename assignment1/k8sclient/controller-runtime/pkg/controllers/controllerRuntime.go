package controllers

import (
	"log"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func ControllerRuntime() client.Client {
	ctrl, err := client.New(config.GetConfigOrDie(), client.Options{})
	if err != nil {
		log.Fatal(err)
	}
	return ctrl

}
