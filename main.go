package main

import (
	"io/ioutil"

	k8sparser "github.com/paoloposso/k8sparser/src"
)

func main() {
	yamlFileBytes, err := ioutil.ReadFile("./tests/examples/mongo.yaml")
	if err != nil {
		panic(err)
	}

	k8sparser.ParseFile(yamlFileBytes)
}
