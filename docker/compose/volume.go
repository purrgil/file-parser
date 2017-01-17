package compose

type ExternalStruct struct {
	Name string
}

type DockerComposeVolume struct {
	Driver     string
	DriverOpts map[string]string
	External   ExternalStruct
	Labels     []string
}

func (dcv *DockerComposeVolume) AddDriverOpt(name string, opt string) {
	if len(dcv.DriverOpts) == 0 {
		dcv.DriverOpts = make(map[string]string)
	}

	dcv.DriverOpts[name] = opt
}

func (dcv *DockerComposeVolume) AddLabel(label string) {
	dcv.Labels = append(dcv.Labels, label)
}

func (dcv *DockerComposeVolume) SetExternal(name string) {
	ex := ExternalStruct{
		Name: name,
	}

	dcv.External = ex
}

func NewVolume() DockerComposeVolume {
	return DockerComposeVolume{
		Driver: "local",
	}
}
