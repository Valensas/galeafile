# Galeafile

Securely manage [Helmfile](https://helmfile.readthedocs.io/en/latest/) releases. Galeafile allows you to make sure
that you are deploying the right environment  to the right namespace and cluster, using a CLI similar to Helmfile.
When an incorrect configuration is detected, the deployment is aborted.

## Installation

### From prebuilt binaries

Navigate to the [Releases](https://github.com/Valensas/galeafile/releases) page and download the binary for your system.

### From source

```bash
git clone git@github.com:Valensas/galeafile.git
cd galeafile
go install .
```

## Configuration

Add `Galeafile.yaml` add where you would like to run Galeafile:

```yaml
# Define the clusters where you deploy your releases
clusters:
  testing:
    servers:
      - https://testing.kubernetes.local
  prod:
    servers:
      - https://prod.kubernetes.local
  local:
    servers:
      - https://localhost:6443
# Define the environments present in your Helmfile
environments:
  # The default environment is used when no environment is specified.
  # Helmfile is run without the -e option.
  default:
    cluster: local
  dev:
    cluster: testing
    # Set the namespace to deploy the releases. If set, runs Helmfile with the -n option
    namespace: dev
    # Set the environment as it appears in Helmfile if it's name is different from the one in Galeafile.
    helmfileEnv: development
  staging:
    cluster: testing
    namespace: staging
  prod:
    cluster: prod

# Set the Helmfile to use. If set, runs Helmfile with the -f option
helmfile: Helmfile.custom.yaml
```

## Usage

- Apply releases: `galeafile apply -e my-env -l name=my-release`
- Diff releases: `galeafile diff -e my-env -l name=my-release`
- Sync releases: `galeafile sync -e my-env -l name=my-release`
- Template releases: `galeafile template -e my-env -l name=my-release`