package purrgil

import (
	"strings"
)

// Package is a struct that map purrgil package object inside purrgil.yml
type Package struct {
	Local         string                 `yaml:"local,omitempty"`
	Download      PurrgilPackageDownload `yaml:"download"`
	Deploy        PurrgilPackageDeploy   `yaml:"deploy,omitempty"`
	PostRunScript []string               `yaml:"post_run, omitempty"`
	LocalScript   map[string]string      `yaml:"local_script,omitempty"`
	DockerScript  map[string]string      `yaml:"docker_script,omitempty"`
}

// SetDownload set a PurrgilPackageDownload into PurrgilPackage
func (pkg *Package) SetDownload(pkgDownload PurrgilPackageDownload) {
	pkg.Download = pkgDownload
}

// AddLocalScript add a script into list of local scripts using a name as a alias for run that local
func (pkg *Package) AddLocalScript(name string, script string) {
	if len(pkg.LocalScript) == 0 {
		pkg.LocalScript = make(map[string]string)
	}

	pkg.LocalScript[name] = script
}

// AddDockerScript add a script into list of docker scripts using a name as a alias for run that into docker container
func (pkg *Package) AddDockerScript(name string, script string) {
	if len(pkg.DockerScript) == 0 {
		pkg.DockerScript = make(map[string]string)
	}

	pkg.DockerScript[name] = script
}

// AddPostRunScript add script to run inside container after the same started
func (pkg *Package) AddPostRunScript(scriptName string) {
	pkg.PostRunScript = append(pkg.PostRunScript, scriptName)
}

// NewPkg create a new PurrgilPackage from a PurrgilPackageDownload
func NewPkg(pkgDownload PurrgilPackageDownload) Package {
	return Package{
		strings.Split(pkgDownload.From, "/")[1],
		pkgDownload,
		PurrgilPackageDeploy{},
		[]string{},
		map[string]string{},
		map[string]string{},
	}
}
