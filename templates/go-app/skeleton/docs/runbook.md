# Runbook

## Pods not ready
Open the **Kubernetes** tab in Backstage and check the pod status and events.

- `ImagePullBackOff`: the image tag in gitops does not exist on Docker Hub. Check the last CI run.
- `CrashLoopBackOff`: check the container logs in the Kubernetes tab.

## Deploy looks stuck
Open the **Argo CD** card on the Overview tab.

- `OutOfSync`: Argo CD has not applied the latest gitops commit yet.
- `Degraded`: the new pods are failing; see "Pods not ready".

## Roll back
Revert the last commit in `gitops/apps/${{ values.name }}`. Argo CD deploys the previous tag.
