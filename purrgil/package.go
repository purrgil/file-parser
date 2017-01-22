package purrgil

type PurrgilPackage struct {
  Local string `yaml:"local,omitempty"`
  Download PurrgilPackageDownload `yaml:"download"`
  Deploy PurrgilPackageDeploy `yaml:"deploy,omitempty"`
  PostRunScript []string `yaml:"post_run, omitempty"`
  LocalScript map[string]string `yaml:"local_script,omitempty"`
  DockerScript map[string]string `yaml:"docker_script,omitempty"`
}

func (pkg *PurrgilPackage) SetDownload (pkgDownload PurrgilPackageDownload) {
  pkg.Download = pkgDownload
}

func (pkg *PurrgilPackage) AddLocalScript(name string, script string) {
  if len(pkg.LocalScript) == 0 {
    pkg.LocalScript = make(map[string]string)
  }

  pkg.LocalScript[name] = script
}

func (pkg *PurrgilPackage) AddDockerScript(name string, script string) {
  if len(pkg.DockerScript) == 0 {
    pkg.DockerScript = make(map[string]string)
  }

  pkg.DockerScript[name] = script
}

func (pkg *PurrgilPackage) AddPostRunScript(scriptName string) {
  pkg.PostRunScript = append(pkg.PostRunScript, scriptName)
}
