# idp-service-templates

Golden paths for Backstage. Each template creates:

1. an app repo (source, Dockerfile, CI, TechDocs, `catalog-info.yaml`, `openapi.yaml`) with CI secrets set
2. a PR on `gitops` adding `apps/<name>/` (Deployment, Service, Kustomization, VirtualService)
3. the catalog entry (Component + API)

Add a template: new folder under `templates/` + one line in `catalog-info.yaml`, same PR.
