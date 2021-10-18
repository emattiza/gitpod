package v1

import (
	"strings"

	"common/cluster"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/xerrors"
)

// Config is the input configuration to create and manage lifecycle of a cluster
//
// Here is a sample representation of yaml configuration
//
// version: v1
// environment: production
// project: gitpod-staging
// #TODO(prs): gcpSecretPath: /var/gcp/gitpod-sa.json
//
// metaClusters:
//   - name: prod-meta-eu00
//     region: europe-west1
//   - name: prod-meta-us01
//     region: us-west-1
// workspaceClusters:
//   - region: europe-west1
//     prefix: eu
//     governedBy: prod-meta-eu01
//     create: true
//     type: gke
//   - region: us-east1
//     prefix: us
//     governedBy: prod-meta-us01
//     create: true
//     type: gke
type Config struct {
	// We do not support cross project deployment
	// All deployments would be in the same GCP project
	Project     *Project
	Version     string              `yaml:"version"`
	Environment cluster.Environment `yaml:"environment"`
	// MetaClusters is optional as we may not want to register the cluster
	MetaClusters      []*cluster.MetaCluster     `yaml:"metaClusters"`
	WorkspaceClusters []cluster.WorkspaceCluster `yaml:"workspaceClusters"`
	// TODO(princerachit): Add gitpod version here when we decide to use installed instead of relying solely on ops repository
}

// Project refers to the project id in GCP
type Project struct {
	Name string `yaml:"name"`
}

func isValidProject(value interface{}) error {
	p, ok := value.(Project)
	if !ok {
		return xerrors.Errorf("value not a valid project")
	}

	err := validation.ValidateStruct(&p,
		validation.Field(strings.TrimSpace(p.Name), validation.Required),
	)
	if err != nil {
		return xerrors.Errorf("invalid project: %s", err)
	}
	return nil
}

// Validate validates if this config is right
func (c *Config) Validate() error {
	err := validation.ValidateStruct(&c,
		validation.Field(c.Version, validation.Required),
		validation.Field(c.Environment, validation.Required),
		validation.Field(c.WorkspaceClusters, validation.Required),
		validation.Field(&c.Project, validation.By(isValidProject)),
		validation.Field(&c.MetaClusters, validation.By(cluster.AreValidMetaClusters)),
		validation.Field(&c.WorkspaceClusters, validation.By(cluster.AreValidWorkspaceClusters)),
	)
	if err != nil {
		return xerrors.Errorf("invalid configuration: %w", err)
	}
	return nil
}

// initializes workspace cluster names based on the config provided
func initializeWorkspaceClusterNames() {

}