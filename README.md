# Strimzi Go APIs

This project contains [Strimzi](https://strimzi.io) APIs for integrating with Strimzi from Go programming language.
It lets you manage the Strimzi resources using the Kubernetes Go Client.

## Updating the APIs

1. Make sure the `doc.go` and `register.go` files are manually created and maintained.
2. Use the `strimzi-go-generate` to generated the Golang-based `types.go`.
3. Go to the `hack` directory and run `codegen.sh` to update the generated files
4. Run `go test` to test the updated files
5. Add the updated files to the Git repo
