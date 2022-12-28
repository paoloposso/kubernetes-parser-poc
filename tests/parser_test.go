package tests

import (
	"io/ioutil"
	"testing"

	k8sparser "github.com/paoloposso/k8sparser/src"
)

func TestParseFile(t *testing.T) {
	yamlFileBytes, err := ioutil.ReadFile("./examples/mongo.yaml")
	if err != nil {
		panic(err)
	}

	err = k8sparser.ParseFile(yamlFileBytes)

	if err != nil {
		t.Error(err)
	}
}
