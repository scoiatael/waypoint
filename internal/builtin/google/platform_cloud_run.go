package google

import (
	"context"

	"github.com/hashicorp/go-hclog"

	"github.com/mitchellh/devflow/internal/builtin/docker"
	"github.com/mitchellh/devflow/internal/datadir"
)

// CloudRunPlatform is the Platform implementation for Google Cloud Run.
type CloudRunPlatform struct {
	config Config
}

// Config implements Configurable
func (p *CloudRunPlatform) Config() interface{} {
	return &p.config
}

// DeployFunc implements component.Platform
func (p *CloudRunPlatform) DeployFunc() interface{} {
	return p.Deploy
}

// Deploy deploys an image to GCR.
func (p *CloudRunPlatform) Deploy(
	ctx context.Context,
	log hclog.Logger,
	img *docker.Image,
	dir *datadir.Component,
) (*CloudRunDeployment, error) {
	log.Warn("DEPLOYING")
	return nil, nil
}

// Config is the configuration structure for the CloudRunPlatform.
type Config struct {
	// Project is the project to deploy to.
	Project string `hcl:"project,attr"`

	// Unauthenticated, if set to true, will allow unauthenticated access
	// to your deployment. This defaults to true.
	Unauthenticated *bool `hcl:"unauthenticated,optional"`
}

// CloudRunDeployment represents a deployment to Google Cloud Run.
type CloudRunDeployment struct{}

// NewCloudRunPlatform is a factory method.
func NewCloudRunPlatform() *CloudRunPlatform { return &CloudRunPlatform{} }
