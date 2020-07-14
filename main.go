package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type KubectlConfig struct {
	Kind           string                    `yaml:"kind"`
	APIVersion     string                    `yaml:"apiVersion"`
	CurrentContext string                    `yaml:"current-context"`
	Clusters       []*KubectlClusterWithName `yaml:"clusters"`
	Contexts       []*KubectlContextWithName `yaml:"contexts"`
	Users          []*KubectlUserWithName    `yaml:"users"`
}

type KubectlClusterWithName struct {
	Name    string         `yaml:"name"`
	Cluster KubectlCluster `yaml:"cluster"`
}

type KubectlCluster struct {
	Server                   string `yaml:"server,omitempty"`
	CertificateAuthorityData []byte `yaml:"certificate-authority-data,omitempty"`
}

type KubectlContext struct {
	Cluster string `yaml:"cluster"`
	User    string `yaml:"user"`
}

type KubectlContextWithName struct {
	Name    string         `yaml:"name"`
	Context KubectlContext `yaml:"context"`
}

type KubectlUserWithName struct {
	Name string      `yaml:"name"`
	User KubectlUser `yaml:"user"`
}

type KubectlUser struct {
	ClientCertificateData []byte `yaml:"client-certificate-data,omitempty"`
	ClientKeyData         []byte `yaml:"client-key-data,omitempty"`
	Password              string `yaml:"password,omitempty"`
	Username              string `yaml:"username,omitempty"`
	Token                 string `yaml:"token,omitempty"`
}

func main() {

	var setting KubectlConfig
	config, err := ioutil.ReadFile("./kubeconfig.yaml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &setting)

	fmt.Println(setting.Kind)
	fmt.Println(setting.APIVersion)
	fmt.Println(setting.CurrentContext)
	fmt.Println(setting.Clusters[0].Cluster.Server)
	fmt.Println(setting.Users[0].User.Token)
	//fmt.Println(setting.SiteNginx.Port)
	//fmt.Println(setting.SiteNginx.LogPath)
	//fmt.Println(setting.SiteNginx.Path)

}
