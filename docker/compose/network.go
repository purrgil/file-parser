package compose

type DockerComposeNetwork struct {
  Driver string
  IPAM Ipam
  External bool
  Labels []map[string]string
}

type Ipam struct {
  Driver string
  Config []IpamConfig
}

type IpamConfig struct {
  SubNet string
  IpRange string
  Gateway string
  AuxAdress []map[string]string
}

func NewNetwork() DockerComposeNetwork {
  return DockerComposeNetwork{}
}
