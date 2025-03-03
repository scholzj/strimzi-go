# Strimzi Go APIs

This project contains [Strimzi](https://strimzi.io) APIs for integrating with Strimzi from Go programming language.
It lets you manage the Strimzi resources using the Kubernetes Go Client.

The following table shows the supported Strimzi versions

| Go API Version | Strimzi version |
|----------------|-----------------|
| `main` branch  | 0.45.0          |
| `0.1.0`        | 0.45.0          |
| `0.2.0`        | 0.45.0          |

## Examples

The [`examples`](./examples) directory contains several examples that show how to use the Strimzi APIs and the generated `ClientSets`.

## Updating the APIs

1. Make sure the `doc.go` and `register.go` files are manually created and maintained.
2. Run `make generate` to update all the files
3. Run `make test` to test the updated files (you need Kubernetes environment with installed Strimzi CRDs)
4. Add the updated files to the Git repo
