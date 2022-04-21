# Microservice template for go-titan
This repository contains template files to generate a base Microservice using Golang and Docker.

## Requirements to create a Microservice with this project:
 - Bash
 - Infuse (https://www.github.com/jucardi/infuse)

## Requirements to run the Microservice once created:
 - Docker and docker-compose
 - Make
 - Golang

## Generating a new Microservice
Within this directory, run the following command:

	$ ./generate [TARGET DIRECTORY]

Where `[TARGET DIRECTORY]` is the directory where the Microservice will be generated.

### Generation configuration

In the root directory of this project, there is a file called `properties.yml`, you may modify this file
to alter how the service will be generated

**Example**

```
service_name: ms-sample
context_path: service
registry: registry.jucardi.io
golang:
  module_path: github.com/jucardi
database: mongo
port: 8080
```

- `service_name` Indicates the name of the service, should match the name of the project in Git
- `context_path` Indicates the base path where the routes will be registered
- `registry` The base URL for the Docker registry where the images will be pushed
- `golang.module_path` The root path to initialize the Golang module. This should not include the service name since it will be automatically appended
- `database` Indicates whether a database will be used for this service. The allowed values are `mongo`, `mysql` and `none`. MySQL currently not supported, comming soon.
- `port` The port where the service will be listening for connections