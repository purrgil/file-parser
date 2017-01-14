package compose

import (
	"gopkg.in/yaml.v2"
)

type DockerCompose struct {
	path     string
	Version  int                             `yaml:"version"`
	Services map[string]DockerComposeService `yaml:"services,omitempty"`
	Volumes  []string                        `yaml:"volumes,omitempty"`
	Network  map[string]DockerComposeNetwork `yaml:"networks,omitempty"`
}

func (dc *DockerCompose) AddService(name string, service DockerComposeService) {
	if len(dc.Services) == 0 {
		dc.Services = make(map[string]DockerComposeService)
	}

	dc.Services[name] = service
}

func (dc *DockerCompose) AddVolume(namedVolume string) {
	dc.Volumes = append(dc.Volumes, namedVolume)
}

func (dc *DockerCompose) AddNetwork(name string, network DockerComposeNetwork) {
	dc.Network[name] = network
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
