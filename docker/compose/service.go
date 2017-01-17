package compose

type ServiceBuild struct {
	Context    string
	Dockerfile string
}

type DockerComposeService struct {
	Build            ServiceBuild      `yaml:"build,omitempty"`
	CapabilitiesAdd  []string          `yaml:"cap_add, omitempty"`
	CapabilitiesDrop []string          `yaml:"cap_drop, omitempty"`
	Command          string            `yaml:"command,omitempty"`
	CGroup           string            `yaml:"cgroup_parent,omitempty"`
	ContainerName    string            `yaml:"container_name,omitempty"`
	Devices          []string          `yaml:"devices,omitempty"`
	DependsOn        []string          `yaml:"depends_on,omitempty"`
	DNS              []string          `yaml:"dns,omitempty"`
	DNSSearch        []string          `yaml:"dns_search,omitempty"`
	TMPFS            []string          `yaml:"tmpfs, omitempty"`
	Entrypoint       string            `yaml:"entrypoint,omitempty"`
	EnvFile          []string          `yaml:"env_file,omitempty"`
	EnvVars          map[string]string `yaml:"environment,omitempty"`
	Image            string            `yaml:"image,omitempty"`
	Links            []string          `yaml:"links,omitempty"`
	Ports            []string          `yaml:"ports,omitempty"`
	Volumes          []string          `yaml:"volumes,omitempty"`
}

/*
	Append a Volume into an array in Service, dcsInstance.Volumes

	Params:

	volume: An string of volume, ex: "./:/app"
*/
func (dcs *DockerComposeService) PushVolume(volume string) {
	dcs.Volumes = append(dcs.Volumes, volume)
}

/*
	Append a Port into an array in Service, dcsInstance.Ports

	Params:

	port: An string of exported ports, ex: "8080:8080"
*/
func (dcs *DockerComposeService) PushPort(port string) {
	dcs.Ports = append(dcs.Ports, port)
}

/*
	Append a Dep into an array in Service, dcsInstance.DependsOn

	Params:

	dep:  string from a service name that want create a dependencie, ex: "db"
*/
func (dcs *DockerComposeService) PushDep(dep string) {
	dcs.DependsOn = append(dcs.DependsOn, dep)
}

/*
	Append a Link into an array in Service, dcsInstance.Links

	Params:
	
	link: a string of service name that you want to link, ex: "db"
*/
func (dcs *DockerComposeService) PushLink(link string) {
	dcs.Links = append(dcs.Links, link)
}

func NewService() DockerComposeService {
	return DockerComposeService{}
}
