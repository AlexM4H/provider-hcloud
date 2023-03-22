package branch

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("hcloud_branch", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "branch"

        // This resource need the repository in which branch would be created
        // as an input. And by defining it as a reference to Repository
        // object, we can build cross resource referencing. See
        // repositoryRef in the example in the Testing section below.
        r.References["repository"] = config.Reference{
            Type: "github.com/AlexM4H/provider-hcloud/apis/repository/v1alpha1.Repository",
        }
    })
}
