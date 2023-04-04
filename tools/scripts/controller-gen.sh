#!/usr/bin/env bash

# Copyright 2022 The RequeueIP Authors.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

# CONST
PROJECT_ROOT=$(dirname ${BASH_SOURCE[0]})/../..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${PROJECT_ROOT}; ls -d -1 ./vendor/sigs.k8s.io/controller-tools/cmd/controller-gen 2>/dev/null || echo ../controller-gen)}

# ENV
# Defines the output path for the artifacts controller-gen generates.
OUTPUT_BASE_PATH=${OUTPUT_BASE_PATH:-${PROJECT_ROOT}/charts/requeueip}
# Diff
CONTROLLER_GEN_TMP_PATH=${CONTROLLER_GEN_TMP_PATH:-${PROJECT_ROOT}/.controller_gen_tmp}
OUTPUT_TMP_PATH=${OUTPUT_TMP_PATH:-${CONTROLLER_GEN_TMP_PATH}/old}
OUTPUT_DIFF_PATH=${OUTPUT_DIFF_PATH:-${CONTROLLER_GEN_TMP_PATH}/new}



controller-gen() {
  go run ${PROJECT_ROOT}/${CODEGEN_PKG}/main.go "$@"
}

manifests_clean() {
  rm -rf ${OUTPUT_BASE_PATH}/crds/*
  rm -rf ${OUTPUT_BASE_PATH}/templates/role.yaml
}

manifests_gen() {
  output_path=$1

  controller-gen \
  crd rbac:roleName="requeueip-admin" \
  paths="${PWD}/${PROJECT_ROOT}/k8s/v1" \
  output:crd:artifacts:config="${output_path}/crds" \
  output:rbac:artifacts:config="${output_path}/templates"
}

deepcopy_gen() {
  trap "cleanup" EXIT SIGINT
  mkdir -p ${CONTROLLER_GEN_TMP_PATH}

  header_file=${CONTROLLER_GEN_TMP_PATH}/boilerplate.go.txt
  cat ${PROJECT_ROOT}/tools/boilerplate.txt | sed -e '$a*/' -e '1i/*' > ${header_file}

  controller-gen \
    object:headerFile="${header_file}" \
    paths="${PWD}/${PROJECT_ROOT}/k8s/v1"
}

manifests_verify() {
  trap "cleanup" EXIT SIGINT

  mkdir -p ${OUTPUT_TMP_PATH}/templates
  if [ "$(ls -A ${OUTPUT_BASE_PATH}/crds)" ]; then
    cp -a ${OUTPUT_BASE_PATH}/crds ${OUTPUT_TMP_PATH}
  fi

  if [ "$(ls -A ${OUTPUT_BASE_PATH}/templates)" ]; then
    cp -a ${OUTPUT_BASE_PATH}/templates/role.yaml ${OUTPUT_TMP_PATH}/templates
  fi

  manifests_gen ${OUTPUT_DIFF_PATH}

  ret=0
  diff -Naupr ${OUTPUT_TMP_PATH} ${OUTPUT_DIFF_PATH} || ret=$?

  if [[ $ret -eq 0 ]];then
    echo "The Artifacts for RequeueIP is up to date"
  else
    echo "Error: The Artifacts for RequeueIP is out of date, run 'make gen.manifests'"
    exit 1
  fi
}

cleanup() {
  rm -rf ${CONTROLLER_GEN_TMP_PATH}
}

help() {
    echo "help"
}

main() {
  case ${1:-none} in
    clean)
      manifests_clean
      ;;
    manifests)
      manifests_clean
      manifests_gen ${OUTPUT_BASE_PATH}
      ;;
    deepcopy)
      deepcopy_gen
      ;;
    verify)
      manifests_verify
      ;;
    *|help|-h|--help)
      help
      ;;
  esac
}

main "$*"