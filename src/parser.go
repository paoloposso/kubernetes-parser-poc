package k8sparser

import (
	"fmt"

	core "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

func Parse(k8sObject []byte) {
	scheme := runtime.NewScheme()
	codecFactory := serializer.NewCodecFactory(scheme)
	deserializer := codecFactory.UniversalDeserializer()

	// namespaceObject, _, err := deserializer.Decode(k8sObject, nil, &core.Namespace{})
	// if err != nil {
	// 	panic(err)
	// }

	// namespace := namespaceObject.(*core.Namespace)
	// fmt.Println("==== namespace.yaml ====")
	// fmt.Printf("Name: %s\n", namespace.ObjectMeta.GetName())
	// fmt.Println()

	// deploymentYAML, err := ioutil.ReadFile("deployment.yaml")
	// if err != nil {
	// 	panic(err)
	// }

	// deploymentObject, _, err := deserializer.Decode(k8sObject, nil, &apps.Deployment{})
	// if err != nil {
	// 	panic(err)
	// }

	// deployment := deploymentObject.(*apps.Deployment)
	// fmt.Println("==== deployment.yaml ====")
	// fmt.Printf("Namespace: %s\n", deployment.ObjectMeta.GetNamespace())
	// fmt.Printf("Resources: %+v\n", deployment.Spec.Template.Spec.Containers[0].Resources.Limits["cpu"])

	serviceObject, _, err := deserializer.Decode(k8sObject, nil, &core.Service{})
	if err != nil {
		panic(err)
	}

	service := serviceObject.(*core.Service)
	fmt.Println("==== deployment.yaml ====")
	fmt.Printf("Resources: %+v\n", service.Name)

}
