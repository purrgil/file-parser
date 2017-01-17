package compose

import (
	"gopkg.in/yaml.v2"
)

/*
	Docker Compose struct to parse YAML to GO Struct
*/
type DockerCompose struct {
	Version  int                             `yaml:"version"`
	Services map[string]DockerComposeService `yaml:"services,omitempty"`
	Volumes  map[string]DockerComposeVolume  `yaml:"volumes,omitempty"`
	Networks map[string]DockerComposeNetwork `yaml:"networks,omitempty"`
}

/*
	This method add a service of type DockerComposeService into .Services -> map[string]DockerComposeService

	You need pass the params:
	
	name: The Service Name, example: app

	service: A DockerComposeService struct instance, you can use NewService() to create one
*/

func (dc *DockerCompose) AddService(name string, service DockerComposeService) {
	if len(dc.Services) == 0 {
		dc.Services = make(map[string]DockerComposeService)
	}

	dc.Services[name] = service
}

/*
	This method add a volume of type DockerComposeVolume into .Volumes -> map[string]DockerComposeVolume

	You need pass the params:

	name: The volume name, example: mysqldata

	volume: A DockerComposeVolume stuct instance, you can use NewVolume()
*/
func (dc *DockerCompose) AddVolume(name string, volume DockerComposeVolume) {
	if len(dc.Volumes) == 0 {
		dc.Volumes = make(map[string]DockerComposeVolume)
	}

	dc.Volumes[name] = volume
}

// You call this AddNetwork(networkName, network)

/*
	This method add a network of type DockerComposeNetwork into .Networks -> map[string]DockerComposeNetwork
	
	You need pass the params:

	name: The network name, example: conn1

	volume: A DockerComposeNetwork stuct instance, you can use NewNetwork()
*/
func (dc *DockerCompose) AddNetwork(name string, network DockerComposeNetwork) {
	if len(dc.Networks) == 0 {
		dc.Networks = make(map[string]DockerComposeNetwork)
	}

	dc.Networks[name] = network
}

/*
	Use "gopkg.in/yaml.v2" to Marshal the Struct
*/
func (dc *DockerCompose) ParseToByte() ([]byte, error) {
	return yaml.Marshal(dc)
}

/*
	Wrap .ParseToByte with string caller to convert byte -> string
*/
func (dc *DockerCompose) ParseToString() (string, error) {
	b, err := dc.ParseToByte()
	return string(b), err
}

/*
	Use "gopkg.in/yaml.v2" to Unmarshal a byte content to a DockerCompose

	You need pass the params:

	content: The load byte from a file, you can use File from this file-parser package to read :)
*/
func ParseToDockerCompose(content []byte) (DockerCompose, error) {
	dc := NewDockerCompose()
	err := yaml.Unmarshal(content, &dc)

	return dc, err
}

/*
	Return a base instance of DockerCompose struct with Version setted as 2
*/
func NewDockerCompose() DockerCompose {
	return DockerCompose{
		Version: 2,
	}
}
