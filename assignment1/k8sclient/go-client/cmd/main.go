package main

import (
	"k8sclient/go-client/pkg"
	"k8sclient/go-client/pkg/controllers"
)

func main() {
	ctrl := controllers.GoClient()
	pkg.ListPods(ctrl)
}
