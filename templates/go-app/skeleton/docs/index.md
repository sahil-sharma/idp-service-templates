# ${{ values.name }}

${{ values.description }}

| | |
|---|---|
| Language | Go |
| Owner | ${{ values.owner }} |
| Image | `docker.io/sahilwork/${{ values.name }}` |
| Deployed by | Argo CD, from `gitops/apps/${{ values.name }}` |

## How a change reaches the cluster

1. Push to `main` in this repo.
2. GitHub Actions builds and pushes the image tagged with the commit SHA.
3. CI updates the image tag in `gitops/apps/${{ values.name }}/kustomization.yaml`.
4. Argo CD syncs the new tag to the cluster.

## Endpoints

| Path | Purpose |
|---|---|
| `/` | Greeting (message from `APP_MESSAGE`) |
| `/healthz` | Liveness and readiness probe |
| `/version` | Running git SHA |
