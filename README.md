# Kubebuilder non-standard directory example

This repo contains a minimal binary (cmd/manager) and a single API type to demonstrate a non standard directory structure for Kubebuilder projects. 

The key insight is that whatever your directory structure, you should update the controller-gen invocations to match.

The most important markers which are added are: 
- group version markers in pkg/api/v1/groupversion_info.go 
    - this would  be required in whatever packaged directly declares your CRD types as go structs.
- object generation markers on the CRD types
- RBAC tags, from the manager binary (or anywhere they might appear)

Other tags like those for webhook, status subresource, also exist.

## Generating CRDs + Manifests

```bash
make generate
```

This should produce a config directory looking like so:
```bash
├── config
│   ├── crd
│   │   └── freeform.alexeldeib.xyz_freeforms.yaml
│   ├── rbac
│   │   └── role.yaml
│   └── webhook
│       └── manifests.yaml
```

Where role will have permissions across secrets due to the annotation in `cmd/manager/main.go`, and the CRD type will be based on the contents of `pkg/api/v1`.
