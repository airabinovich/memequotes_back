#!/usr/bin/env bash

PROJECT_DIR=$(dirname "$0")

cd "${PROJECT_DIR}"
[[ -r "coverage.out"  ]] && rm "coverage.out"
go test ./... -race -coverprofile=coverage.out && go tool cover -html=coverage.out
cd -

