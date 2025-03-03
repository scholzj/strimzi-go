#!/bin/bash

##########
# Inspired by https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/hack/update-codegen.sh
##########

set -o errexit
set -o nounset
set -o pipefail

go get k8s.io/code-generator

GO_CMD=${1:-go}
CURRENT_DIR=$(dirname "${BASH_SOURCE[0]}")
CODEGEN_PKG=$($GO_CMD list -m -mod=readonly -f "{{.Dir}}" k8s.io/code-generator)
cd "${CURRENT_DIR}/.."

# shellcheck source=/dev/null
source "${CODEGEN_PKG}/kube_codegen.sh"

echo "Generating helpers..."

kube::codegen::gen_helpers \
    --boilerplate ./hack/boilerplate.go.txt \
    ./pkg/apis/kafka.strimzi.io

kube::codegen::gen_helpers \
    --boilerplate ./hack/boilerplate.go.txt \
    ./pkg/apis/core.strimzi.io

echo "Generating client code..."

kube::codegen::gen_client \
  ./pkg/apis \
  --output-pkg github.com/scholzj/strimzi-go/pkg/client \
  --output-dir ./pkg/client \
  --boilerplate hack/boilerplate.go.txt \
  --with-watch

echo "Running go mod tidy ..."

# We need to clean up the go.mod file since code-generator adds temporary library to the go.mod file.
"${GO_CMD}" mod tidy