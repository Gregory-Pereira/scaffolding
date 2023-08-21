package build

import (
	"os"
)

func GetBuildMethod() string {
	kuberenetesServiceHost := os.Getenv("KUBERNETES_SERVICE_HOST")
	if kuberenetesServiceHost == "" {
		return "container"
	} else {
		return "kuberenetes"
	}
}
