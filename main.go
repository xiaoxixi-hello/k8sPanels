package main

import "github.com/ylinyang/k8sPanels/router"

func main() {
	r := router.Router()
	if err := r.Run(":8888"); err != nil {
		panic(err)
	}
}
