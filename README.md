# Gotter

[![CI](https://github.com/josephrodriguez/gotter/actions/workflows/go-build.yaml/badge.svg)](https://github.com/josephrodriguez/gotter/actions)
[![Docker Pulls](https://img.shields.io/docker/pulls/josephrodriguez/gotter?logo=docker)](https://hub.docker.com/r/josephrodriguez/gotter)
[![Go Version](https://img.shields.io/badge/go-1.25%2B-blue?logo=go)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

![Gotter Logo](./docs/logo.png)

Gotter keeps your GitOps repository up-to-date by automatically syncing Helm chart versions (from Artifactory and other registries) into Kustomize manifests.


Configuration:

| Variable | Description | Default |
|----------|-------------|---------|
| `GOTTER_HOST` | Server host address | `0.0.0.0` |
| `GOTTER_PORT` | Server port | `8080` |


Supported Registries & Manifests

| Registry     | Supported Manifests | Integration Type | Notes |
|--------------|---------------------|------------------|-------|
| Artifactory  | Kustomize           | Webhook + API    | Detects Helm chart publishes and updates Kustomize manifests (patches kustomization YAML or patch files) |

Roadmap

- Add support for additional registries (ChartMuseum, Harbor, OCI registries) and manifest formats (Helm values, plain YAML) in future releases.

See the full README in the repository for detailed usage and contributing guidelines.