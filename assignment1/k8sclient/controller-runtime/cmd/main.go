package main

import (
	"k8sclient/controller-runtime/pkg"
	"k8sclient/controller-runtime/pkg/controllers"
)

func main() {
	ctrl := controllers.ControllerRuntime()
	pkg.ListPods(ctrl)
}
