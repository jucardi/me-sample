#!/bin/sh

go mod init {{.golang.module_path}}/{{.service_name}}
go mod tidy
go mod vendor
rm init.sh
rm generate.sh
rm README.md
mv README.template.md README.md