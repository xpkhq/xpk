# Concept

`xpk` is not only a package manager.

It is made up of four core concepts.

## Fetch

`xpk` can download packages in `.xpk` format from an xpk source (e.g. xpkhub).

## Build

`.xpk` packages may contain pre-built filesystems, but sometimes it will need to be compiled.

`xpk` can run compile process automatically, generating the whole file tree.

## Deploy

As soon as there is a pre-built filesystem, it can be deployed.

There are two types of deployment:

### Apply to host machine

This will create symbol links in the host machine.

Usually, packages in `utils` category will be deployed using this method.

### Run service

This will create a container and runs the service in it.

A package can be deployed many times through this way.

As a result, there are three types of dependencies:

- Build dependency: required when need to build the package from bootstrap
- Runtime dependency: always required
- Service dependency: required when deploying a service. Usually another running service.

## Compose

`xpk` can compose between different machines.
