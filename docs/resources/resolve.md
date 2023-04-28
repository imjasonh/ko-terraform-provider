---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ko_resolve Resource - terraform-provider-ko"
subcategory: ""
description: |-
  
---

# ko_resolve (Resource)



## Example Usage

```terraform
provider "ko" {
  repo = "ttl.sh/tf-ko"
}

resource "ko_build" "example" {
  importpath = "github.com/google/ko/test"
}

output "image_ref" {
  value = ko_build.example.image_ref
}

resource "ko_resolve" "example" {
  filenames = ["../../../testdata/k8s.yaml"]
  recursive = false
}

output "manifests" {
  value = ko_resolve.example.manifests
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `filenames` (List of String) Filenames, directories, or URLs to files to use to create the resource

### Optional

- `base_image` (String)
- `platforms` (List of String) Platforms to build for, comma separated. e.g. linux/amd64,linux/arm64
- `push` (Boolean) Push images to KO_DOCKER_REPO
- `recursive` (Boolean) Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.
- `sbom` (String) The SBOM media type to use (none will disable SBOM synthesis and upload, also supports: spdx, cyclonedx, go.version-m).
- `selector` (String) Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)
- `tags` (List of String) Tags to apply to the image, comma separated. e.g. latest,1.0.0
- `working_dir` (String) The working directory to use for the build context.

### Read-Only

- `id` (String) The ID of the resource.
- `manifests` (List of String) The manifests created by the resource.

