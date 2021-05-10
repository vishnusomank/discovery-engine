package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetK8sNamespaces(t *testing.T) {
	actual := GetNamespacesK8s()

	assert.Contains(t, actual, "kube-system")
}

func TestGetConGroups(t *testing.T) {
	actual := GetPodsK8s()

	for _, pod := range actual {
		if pod.PodName == "" {
			t.Errorf("it should have a pod name")
		}
	}
}

func TestGetServices(t *testing.T) {
	actual := GetServicesK8s()

	for _, svc := range actual {
		if svc.ServiceName == "" {
			t.Errorf("it should have a service name")
		}
	}
}

func TestGetEndpoints(t *testing.T) {
	actual := GetEndpointsK8s()

	for _, endpoint := range actual {
		if endpoint.EndpointName == "" {
			t.Errorf("it should have a endpoint name")
		}
	}
}
