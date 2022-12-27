package main

import (
	"io/ioutil"

	k8sparser "github.com/paoloposso/k8sparser/src"
)

func main() {
	k8sObject, err := ioutil.ReadFile("../mongo-demo/srv1.yaml")
	if err != nil {
		panic(err)
	}
	k8sparser.Parse(k8sObject)
}
