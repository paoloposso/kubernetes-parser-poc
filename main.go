package main

import (
	"fmt"
	"io/ioutil"

	k8sparser "github.com/paoloposso/k8sparser/src"
)

func main() {

	files, _ := ioutil.ReadDir("./tests/examples")

	fmt.Println(files)

	for _, item := range files {
		fmt.Println(item.Name())
		yamlFileBytes, err := ioutil.ReadFile(fmt.Sprintf("./tests/examples/%s", item.Name()))

		if err != nil {
			panic(err)
		}

		k8sparser.ParseFile(yamlFileBytes)
	}
}
