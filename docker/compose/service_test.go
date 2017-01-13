package compose

import (
	"reflect"
	"testing"
)

func TestPushFunctions(t *testing.T) {
	baseRep := []string{"test"}
	finalRep := append(baseRep, "x")

	dcs := DockerComposeService{
		CapabilitiesAdd:  baseRep,
		CapabilitiesDrop: baseRep,
		Devices:          baseRep,
		DependsOn:        baseRep,
		DNS:              baseRep,
		DNSSearch:        baseRep,
		TMPFS:            baseRep,
		EnvFile:          baseRep,
		Links:            baseRep,
		Ports:            baseRep,
		Volumes:          baseRep,
	}

	for _, test := range []struct {
		method   string
		field    string
		response []string
	}{
		{"PushVolume", "Volumes", finalRep},
		{"PushPort", "Ports", finalRep},
		{"PushDep", "DependsOn", finalRep},
		{"PushLink", "Links", finalRep},
	} {
		ref := reflect.ValueOf(&dcs)
		method := ref.MethodByName(test.method)

		in := reflect.ValueOf("x")
		method.Call([]reflect.Value{in})

		fieldValue := ref.Elem().FieldByName(test.field).Interface()
		if !reflect.DeepEqual(fieldValue.([]string), finalRep) {
			t.Error(
				"Expected", finalRep,
				"got", fieldValue,
			)
		}

	}
}
