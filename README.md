# Strimzi Go APIs

This project contains [Strimzi](https://strimzi.io) APIs for integrating with Strimzi from Go programming language.
It lets you manage the Strimzi resources using the Kubernetes Go Client.

The following table shows the supported Strimzi versions

| Go API Version | Strimzi version |
|----------------|-----------------|
| `main` branch  | 0.45.0          |

## TODO:

* [ ] Add support for `StrimziPodSet` resources
* [ ] Add build automation
* [ ] First release

## Examples

The [`examples`](./examples) directory contains several examples that show how to use the Strimzi APIs and the generated `ClientSets`.

## Updating the APIs

1. Make sure the `doc.go` and `register.go` files are manually created and maintained.
2. Use the `strimzi-go-generate` to generated the Golang-based `types.go`.
3. Go to the `hack` directory and run `codegen.sh` to update the generated files
4. Run `go test ./tests/...` to test the updated files
5. Add the updated files to the Git repo
