#!/usr/bin/env bash

BASEDIR=$(dirname "$0")
OUTDIR=$(realpath "${BASEDIR}/../protos")
#OUTDIR_JS=$(realpath "${BASEDIR}/../../protosjs/homework")

pushd "${BASEDIR}" || (echo "failed to change directory to ${BASEDIR}" && exit)

#build
#protoc --proto_path=../validate --proto_path=. --go_out=plugins=grpc:"${OUTDIR}" ./*.proto --validate_out="lang=go:${OUTDIR}"
protoc --go_out=${OUTDIR} --go_opt=paths=source_relative \
       --go-grpc_out=${OUTDIR} --go-grpc_opt=paths=source_relative \
       ./*.proto

#build to js
#protoc --proto_path=../validate --proto_path=. --js_out=import_style=commonjs,binary:"${OUTDIR_JS}" ./*.proto
