package compose

import (
	"testing"
)

func TestNewVolume(t *testing.T) {
	dcv := NewVolume()

	if dcv.Driver != "local" {
		t.Error(
			"Expected", "local",
			"got", dcv.Driver,
		)
	}
}

func TestAddDriverOpt(t *testing.T) {
	driverOptName := "test_name"
	driverOptValue := "test_value"

	dcv := NewVolume()
	dcv.AddDriverOpt(driverOptName, driverOptValue)

	if dcv.DriverOpts[driverOptName] != driverOptValue {
		t.Error(
			"Expected", driverOptValue,
			"got", dcv.DriverOpts[driverOptName],
		)
	}
}

func TestAddLabel(t *testing.T) {
	label := "label:test"

	dcv := NewVolume()
	dcv.AddLabel(label)

	if len(dcv.Labels) != 1 {
		t.Error(
			"Expected", "len == 1",
			"got", len(dcv.Labels),
		)
	}
}

func TestSetExternal(t *testing.T) {
	externalName := "test"

	dcv := NewVolume()
	dcv.SetExternal(externalName)

	if dcv.External.Name != externalName {
		t.Error(
			"Expected", externalName,
			"got", dcv.External.Name,
		)
	}
}
