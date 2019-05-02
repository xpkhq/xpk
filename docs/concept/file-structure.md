# File structure

The structure of `.xpk` file is defined here.

## Overall

First, it's a `.tar.gz` archive.

> Why?
>
> Just because it's easier to de-compress in Golang.

## Then

In the root directory, there are two directories and a file.

The file is xpk.yml, including basic information of this package.

example:

```yaml
# xpk.yml: basic information of the package

# Name of the package
name: testPkg
# A short description
description: testPkg is a package providing some useless tools.
# Type, can be [util|service]
type: util
# Authors
author:
  - name: testPerson
    email: test@example.com
  - name: testPM
    email: project-manager@are-dogs.com
# Dependencies
dependencies:
  build:
    gcc: ">=4.8.0"
    make: ">=2.3.4"
  runtime:
    binutils: ">=0.1.0"
```

The two directories are `build` and `fs`.

In `build` directory, there are source codes and other build scripts.

In `fs` directory, there are more directories named platforms.

For example:

```tree
.
├── linux
│   ├── amd64
│   ├── armv6
│   ├── armv7
│   ├── i386
│   ├── mips
│   └── x86_64
├── osx
│   └── x86_64
└── windows
    ├── i386
    └── x86_64
```
