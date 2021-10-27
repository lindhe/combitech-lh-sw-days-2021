# GitOps with Fleet

_Are you still stuck with Continuous Integration?_

## `fleet-install/`

The `fleet-install` directory holds the configuration needed to install Fleet in a Kubernetes cluster.
While there are some Helm charts available via GitHub, they are not yet
available via a public Helm repo, so we package it as a script and may use e.g.
[Helmfile](https://github.com/roboll/helmfile) to configure it in the future.
