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

func (dcs *DockerComposeService) PushVolume(volume string) {
	dcs.Volumes = append(dcs.Volumes, volume)
}

func (dcs *DockerComposeService) PushPort(port string) {
	dcs.Ports = append(dcs.Ports, port)
}

func (dcs *DockerComposeService) PushDep(dep string) {
	dcs.DependsOn = append(dcs.DependsOn, dep)
}

func (dcs *DockerComposeService) PushLink(link string) {
	dcs.Links = append(dcs.Links, link)
}

func NewService() DockerComposeService {
	return DockerComposeService{}
}
