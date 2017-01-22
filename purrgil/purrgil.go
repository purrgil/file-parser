package purrgil

type Purrgil struct {
  Name string `yaml:"name"`
  Packages map[string]PurrgilPackage `yaml:"packages, omitempty"`
}

func (p *Purrgil) AddPackage(name string, pkg PurrgilPackage) {
    if len(p.Packages) == 0 {
      p.Packages = make(map[string]PurrgilPackage)
    }

    p.Packages[name] = pkg
}
