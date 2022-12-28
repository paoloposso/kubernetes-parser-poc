package k8sparser

import (
	"errors"
	"fmt"
	"strings"

	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"

	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

type genericK8sObject struct {
	Kind string `json:"kind,omitempty"`
}

func ParseFile(fileBytes []byte) error {
	yamlFileString := string(fileBytes)

	yamlSectionsList := strings.Split(yamlFileString, "---")

	for _, elem := range yamlSectionsList {
		err := parseK8sElement([]byte(elem))
		if err != nil {
			return err
		}
	}
	return nil
}

func parseK8sElement(k8sObject []byte) error {
	scheme := runtime.NewScheme()
	codecFactory := serializer.NewCodecFactory(scheme)

	var k8sMeta genericK8sObject
	err := yaml.Unmarshal([]byte(k8sObject), &k8sMeta)

	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding YAML object: %s", err))
	}

	deserializer := codecFactory.UniversalDeserializer()

	if k8sMeta.Kind == "Deployment" {
		deploymentObject, _, err := deserializer.Decode(k8sObject, nil, &apps.Deployment{})
		if err != nil {
			panic(err)
		}
		deployment := deploymentObject.(*apps.Deployment)
		fmt.Println("==== deployment.yaml ====")
		fmt.Printf("Namespace: %s\n", deployment.ObjectMeta.GetNamespace())
		fmt.Printf("Resources: %+v\n", deployment.Spec.Template.Spec.Containers[0].Resources.Limits["cpu"])
	} else if k8sMeta.Kind == "Service" {
		obj, _, err := deserializer.Decode(k8sObject, nil, &core.Service{})
		if err != nil {
			panic(err)
		}
		service := obj.(*core.Service)
		fmt.Println("==== deployment.yaml ====")
		fmt.Printf("Namespace: %s\n", service.ObjectMeta.GetNamespace())
		fmt.Printf("ClusterIP: %+v\n", service.Spec.ClusterIP)
	}

	return nil
}
