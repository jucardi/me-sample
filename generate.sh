#!/bin/bash

function __print_usage() {
    echo "
This directory contains a sample template for a Golang microservice. To generate a new
Microservice, modify the 'properties.yml' file to properly specify the parameters of
your service. Once saved, run the 'generate' formula in this Makefile

USAGE:  ./generate.sh '/path/to/destination'

NOTE: The command 'infuse' is required to generate the templates, to install it, run
      './generate deps' Golang is required to install infuse. For more information about
      infuse, visit https://www.github.com/jucardi/infuse
"
}

if [ "$1" == "deps" ]; then
    go install github.com/jucardi/infuse/cmd/infuse@latest
    exit 0
elif [ "$1" == "usage" ]; then
    __print_usage
    exit 0
elif [ "$1" != "" ]; then
    TARGET="$1"
fi

if [ "${TARGET}" == "" ]; then
    echo "Target directory is required"
    __print_usage
     exit 1
fi

echo "
Generating Microservice at destination: ${TARGET}

Using the following configuration
"
cat properties.yml
echo ""
infuse . -f properties.yml -o "${TARGET}" --ignoreErrors

cd "${TARGET}"
chmod +x init.sh
./init.sh

