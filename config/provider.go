/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

        "github.com/AlexM4H/provider-hcloud/config/branch"
        "github.com/AlexM4H/provider-hcloud/config/repository"
)

const (
	resourcePrefix = "hcloud"
	modulePath     = "github.com/AlexM4H/provider-hcloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
                repository.Configure,
                branch.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
