package models

type MavenLibrary struct {
	Coordinates string    `json:"coordinates,omitempty" url:"coordinates,omitempty"`
	Repo        string    `json:"repo,omitempty" url:"repo,omitempty"`
	Exclusions  *[]string `json:"exclusions,omitempty" url:"exclusions,omitempty"`
}
