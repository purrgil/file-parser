package purrgil

import (
	"gopkg.in/yaml.v2"
)

// Purrgil is a simple struct that defines Purrgil format
type Purrgil struct {
	Name     string             `yaml:"name"`
	Packages map[string]Package `yaml:"packages,omitempty"`
}

// New create a New instance o Purrgil Struct
func New(name string) Purrgil {
	return Purrgil{
		name,
		map[string]Package{},
	}
}

// AddPackage add a Package inside purrgil configuration
//	@param name <string> The package name
//	@param pkg <Package> The data from package
func (p *Purrgil) AddPackage(name string, pkg Package) {
	if len(p.Packages) == 0 {
		p.Packages = make(map[string]Package)
	}

	p.Packages[name] = pkg
}

// ParseToByte Marshall all Purrgil configuration to []byte
func (p *Purrgil) ParseToByte() ([]byte, error) {
	return yaml.Marshal(p)
}

// ParseToString Marshall all Purrgil configuration to string
func (p *Purrgil) ParseToString() (string, error) {
	b, err := p.ParseToByte()
	return string(b), err
}

// ParseToPurrgil Unmarshal []bytes to a Purrgil Struct
func ParseToPurrgil(content []byte) (Purrgil, error) {
	p := New("")
	err := yaml.Unmarshal(content, &p)

	return p, err
}
