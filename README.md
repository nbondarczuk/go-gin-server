# Go Gin Server Building Toolkit

## Purpose

The project uses a basic template of a Golang Gin HTTP server to build
a generic set of make file rules. They are usually less standardized
then Golang web servers. The build rules facilitate Golang compilation,
docker image creation and running docker compose local integration tests.

## Silver bullets

### Gin server template

The server provides the most simple access point: /health. It has basic
configuration. It uses standard kibrary package slog.

### VCS freedom

The project is VCS independent as it uses replace statement in go.mod.

Vide: https://go.dev/wiki/Modules#can-i-work-entirely-outside-of-vcs-on-my-local-filesystem

### Make build rules

The rules are stored in the build/include directory. The main Makefile
includes all files from this path. This makes sense as the files may
added, modified, removed and there is not fixed list of them.

### Docker image building

The provided example of a Dockerfile is a pretty generic template. It shows
how to build docker image in 2-stage process.

### Run testing

The curl script is used to test basic server run in the local box.

### Integration testing

This is docker used in docker compose setup. Just simple test of the functionality
is done here.

### Performance testing

This is more complex case of docker compose as a testing client like ab starts
to run the server API access points.

### Kubernetes deployment

The k8s deployment is done with minikube/* target. It creates 3 node service using minikube.
