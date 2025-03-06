package main

import (
	"github.com/m4xkub/capstonev2_worker/controller"
)

func main() {
	r := controller.GetRootController()
	r.Run(":8080")
}
