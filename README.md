# GitOps with Fleet

_Are you still stuck with Continuous Integration?_

## `fleet-examples/`

In the `fleet-examples` directory, we have forked the [official example repo for Fleet](https://github.com/rancher/fleet-examples/).
This is mostly just to have some sane samples to test to make sure everything works as intended.

The directory is a `git subtree`.
To sync it with the upstream repo, run the following command:

```console
git subtree pull --prefix fleet-examples --squash https://github.com/rancher/fleet-examples.git master
```
