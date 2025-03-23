#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# corresponding to go mod init <module>
MODULE=controller
# apis package
APIS_PKG=pkg/apis
# generated output package
OUTPUT_PKG=internal
# group-version
GROUP=cloudnative.group
VERSION=v1
GROUP_VERSION=${GROUP}:${VERSION}

#echo $(dirname ${BASH_SOURCE[0]};pwd)
#exit 1

#SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
SCRIPT_ROOT=$(dirname $(readlink -f "$0"))/..
echo ${SCRIPT_ROOT}

cd "${SCRIPT_ROOT}";

CODEGEN_PKG=$(ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)

#CODEGEN_PKG=$(CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator))

rm -rf ${OUTPUT_PKG}/{clientset,informers,listers}

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
#bash "${CODEGEN_PKG}"/generate-groups.sh "client,informer,lister" \
bash "${CODEGEN_PKG}"/generate-groups.sh all \
  controller/${OUTPUT_PKG} \
  controller/${APIS_PKG} \
  ${GROUP_VERSION} \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --output-base "${SCRIPT_ROOT}" \
  -v=1

cp -rf $SCRIPT_ROOT/controller/${OUTPUT_PKG} $SCRIPT_ROOT
cp -rf $SCRIPT_ROOT/controller/${APIS_PKG} $SCRIPT_ROOT/pkg
rm -rf $SCRIPT_ROOT/controller/