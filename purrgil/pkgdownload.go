package purrgil

type PurrgilPackageDownload struct {
  Provider string `yaml:"provider"`
  From string `yaml:"from"`
  SSH bool `yaml:"ssh"`
  PostInstall []string `yaml:"post_install"`
}
