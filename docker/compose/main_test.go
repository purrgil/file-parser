package compose

import (
	"reflect"
	"testing"
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
			"but got", dc.Volumes[testName],
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

func TestParseToString(t *testing.T) {
	dc := NewDockerCompose()

	b, _ := dc.ParseToByte()
	s, _ := dc.ParseToString()

	if string(b) != s {
		t.Error(
			"Expected", string(b),
			"got", s,
		)
	}
}

func TestParseToDockerCompose(t *testing.T) {
	myByteDocker := []byte(`
		version: "2"
		services:
			test:
				build:
					context: ./
					dockerfile: Dockerfile.dev
	`)

	dc, _ := ParseToDockerCompose(myByteDocker)

	if dc.Services["test"].Build.Context == "./" {
		t.Error(
			"Expected service[test]", "./",
			"got", dc.Services["test"].Build.Context,
		)
	}
}
