package models

type Library struct {
	Jar   string            `json:"jar,omitempty" url:"jar,omitempty"`
	Egg   string            `json:"egg,omitempty" url:"egg,omitempty"`
	Whl   string            `json:"whl,omitempty" url:"whl,omitempty"`
	Pypi  PythonPyPiLibrary `json:"pypi,omitempty" url:"pypi,omitempty"`
	Maven MavenLibrary      `json:"maven,omitempty" url:"maven,omitempty"`
	Cran  RCranLibrary      `json:"cran,omitempty" url:"cran,omitempty"`
}
