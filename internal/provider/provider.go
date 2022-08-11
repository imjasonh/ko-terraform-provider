package provider

import (
	"context"
	"fmt"

	"github.com/google/ko/pkg/commands/options"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var changed = false

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"docker_repo": {
					Description: "Container repositor to publish images to. Defaults to `KO_DOCKER_REPO` env var",
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KO_DOCKER_REPO", ""),
					Type:        schema.TypeString,
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"ko_image":   resourceImage(),
				"ko_resolve": resolveConfig(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

// configure initializes the global provider with sensible defaults (that mimic what ko does with cli/cobra defaults)
func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, s *schema.ResourceData) (interface{}, diag.Diagnostics) {
		repo, ok := s.Get("docker_repo").(string)
		if !ok {
			return nil, diag.Errorf("expected docker_repo to be string")
		}
		if repo == "" {
			return nil, diag.Errorf("docker_repo attribute or KO_DOCKER_REPO environment variable must be set")
		}

		return &providerOpts{
			bo: &options.BuildOptions{},
			po: &options.PublishOptions{
				DockerRepo: repo,
			},
		}, nil
	}
}

type providerOpts struct {
	bo *options.BuildOptions
	po *options.PublishOptions
}

func NewProviderOpts(meta interface{}) (*providerOpts, error) {
	opts, ok := meta.(*providerOpts)
	if !ok {
		return nil, fmt.Errorf("parsing provider args: %v", meta)
	}

	// This won't parse the cmd flags, but it will parse any environment vars and set some helpful defaults for us
	if err := opts.bo.LoadConfig(); err != nil {
		return nil, err
	}

	return opts, nil
}
