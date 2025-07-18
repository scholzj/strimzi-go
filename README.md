# Strimzi Go APIs

This project contains [Strimzi](https://strimzi.io) APIs for integrating with Strimzi from Go programming language.
It lets you manage the Strimzi resources using the Kubernetes Go Client.

The following table shows the supported Strimzi versions

| Go API Version | Strimzi version |
|----------------|-----------------|
| `main` branch  | 0.45.0          |
| `0.1.0`        | 0.45.0          |
| `0.2.x`        | 0.45.0          |
| `0.3.x`        | 0.46.0          |
| `0.4.x`        | 0.47.x          |

## Examples

The [`examples`](./examples) directory contains several examples that show how to use the Strimzi APIs and the generated `ClientSets`.

## Users

If you are using the Strimzi Go API, feel free to open a PR and add your project here so that others can see what you created and use it as an example if needed.

* [Keksposé](https://github.com/scholzj/kekspose): Expose your Strimzi-based Apache Kafka cluster outside your Minikube, Kind, or Docker Desktop clusters

## Updating the APIs

To add support for new Strimzi version, you should try to follow these steps:

1. Update the Strimzi version in the `pom.xml` file of the `strimzi-go-generator`.
2. Update the Strimzi CRD version installed in the GitHub Actions (in the `build.yaml` file).
3. Make sure the `doc.go` and `register.go` files are manually created and maintained.
4. If needed, update the Kubernetes versions in `go.mod` (e.g. `k8s.io/api` and `k8s.io/client-go`).
   The generator will always generate the code based on the on the latest Kube version so not updating them might cause error.
5. Run `make generate` to update all the generated files. 
   This runs both the Java generator to generate the CRD types as well the Kubernetes client generator to generate the Go-lang client code.
6. Run `make build` to build the Go project files.
7. Run `make test` to test the updated files (you need Kubernetes environment with installed Strimzi CRDs).
8. Update the `README.md` file to indicate the new supported Strimzi version.
9. Add the updated files to the Git repo.
