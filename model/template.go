package model

import "github.com/rancher/go-rancher/client"

type Template struct {
	client.Resource
	Name           string            `json:"name"`
	UUID 		   string			 `json:"uuid"`
	Category       string            `json:"category"`
	Description    string            `json:"description"`
	Version        string            `json:"version"`
	IconLink       string            `json:"iconLink"`
	VersionLinks   map[string]string `json:"versionLinks"`
	DockerCompose  string            `json:"dockerCompose"`
	RancherCompose string            `json:"rancherCompose"`
	Questions      []Question        `json:"questions"`
	Path           string            `json:"path"`
}

type TemplateCollection struct {
	client.Collection
	Data []Template `json:"data,omitempty"`
}

/*var Schemas = Schemas{
	Data: []Schema{
		{},
		{},
	},
}*/
