#!/bin/bash

# setup
folder="generated"
package="generated"
oapiFile="petstore-expanded.yaml"

# cleanup
rm -rf ./$folder
mkdir -p ./$folder

# codegen
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
oapi-codegen -generate types -o $folder/$package-types.gen.go -package $package $oapiFile
oapi-codegen -generate "chi-server" -o $folder/$package.gen.go -package $package $oapiFile

# go
go mod tidy