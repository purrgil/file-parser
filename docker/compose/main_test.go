package compose

import (
	"testing"
	"reflect"
)

const testName string = "test_name"

func TestAddNetwork(t *testing.T) {
	dc := NewDockerCompose()
	testNetwork := NewNetwork()

	dc.AddNetwork(testName, testNetwork)

	if !reflect.DeepEqual(dc.Networks[testName], testNetwork) {
		t.Error(
			"Expected network", testNetwork,
			"but got", dc.Networks[testName],
		)
	}
}

func TestAddVolume(t *testing.T) {
	dc := NewDockerCompose()
	testVolume := NewVolume()

	dc.AddVolume(testName, testVolume)

	if !reflect.DeepEqual(dc.Volumes[testName], testVolume) {
		t.Error(
			"Expected volumes", testVolume,
			"but got",	dc.Volumes[testName],
		)
	}
}

func TestAddService(t *testing.T) {
	dc := NewDockerCompose()
	testService := NewService()

	dc.AddService(testName, testService)

	if !reflect.DeepEqual(dc.Services[testName], testService) {
		t.Error(
			"Expected services", testService,
			"but got", dc.Services[testName],
		)
	}
}
