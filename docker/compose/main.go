package compose

import (
	"gopkg.in/yaml.v2"
)

type DockerCompose struct {
	path     string
	Version  int                             `yaml:"version"`
	Services map[string]DockerComposeService `yaml:"services,omitempty"`
	Volumes  map[string]DockerComposeVolume  `yaml:"volumes,omitempty"`
	Networks map[string]DockerComposeNetwork `yaml:"networks,omitempty"`
}

func (dc *DockerCompose) AddService(name string, service DockerComposeService) {
	if len(dc.Services) == 0 {
		dc.Services = make(map[string]DockerComposeService)
	}

	dc.Services[name] = service
}

func (dc *DockerCompose) AddVolume(name string, volume DockerComposeVolume) {
	if len(dc.Volumes) == 0 {
		dc.Volumes = make(map[string]DockerComposeVolume)
	}

	dc.Volumes[name] = volume
}

func (dc *DockerCompose) AddNetwork(name string, network DockerComposeNetwork) {
	if len(dc.Networks) == 0 {
		dc.Networks = make(map[string]DockerComposeNetwork)
	}

	dc.Networks[name] = network
}

func (dc *DockerCompose) ParseToByte() ([]byte, error) {
	return yaml.Marshal(dc)
}

func (dc *DockerCompose) ParseToString() (string, error) {
	b, err := dc.ParseToByte()
	return string(b), err
}

func ParseToDockerCompose(content []byte) (DockerCompose, error) {
	dc := NewDockerCompose()
	err := yaml.Unmarshal(content, &dc)

	return dc, err
}

func NewDockerCompose() DockerCompose {
	return DockerCompose{
		Version: 2,
	}
}
