package model

import "fmt"

func assembleNamespace(namespace string, name string) string {
	return fmt.Sprintf("%s%s", namespace, name[1:])
}

func assembleServiceName(name string) string {
	return fmt.Sprintf("%sService", name)
}
