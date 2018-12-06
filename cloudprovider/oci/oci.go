package oci

import (
	"encoding/json"
	"fmt"

	"github.com/rancher/types/apis/management.cattle.io/v3"
)

const (
	OciCloudProviderName = "oci"
)

type CloudProvider struct {
	Config *v3.OciCloudProvider
	Name   string
}

func GetInstance() *CloudProvider {
	return &CloudProvider{}
}

func (p *CloudProvider) Init(cloudProviderConfig v3.CloudProvider) error {
	if cloudProviderConfig.OciCloudProvider == nil {
		return fmt.Errorf("Oci Cloud Provider Config is empty")
	}
	p.Name = OciCloudProviderName
	if cloudProviderConfig.Name != "" {
		p.Name = cloudProviderConfig.Name
	}
	p.Config = cloudProviderConfig.OciCloudProvider
	return nil
}

func (p *CloudProvider) GetName() string {
	return p.Name
}

func (p *CloudProvider) GenerateCloudConfigFile() (string, error) {
	cloudConfig, err := json.MarshalIndent(p.Config, "", "\n")
	if err != nil {
		return "", err
	}
	return string(cloudConfig), nil
}
