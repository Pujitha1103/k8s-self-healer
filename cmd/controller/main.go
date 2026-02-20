package main

import (
    "fmt"

    "github.com/pujithapotta/k8s-self-healer/pkg/controller"
)

func main() {
    fmt.Println("k8s-self-healer starting")
    c := controller.NewController()
    c.Run()
}
