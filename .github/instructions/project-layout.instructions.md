This repository SHOULD keep simple layout.

# Go Module Layout

Referring to https://go.dev/doc/modules/layout, a Toolkit module can split into
differenct Exported package under Module root directory.

For this tool library repository, you only need to organize the code based on different logical module packages and build unit tests within those packages.

No `main` package code is required, nor `cmd/*` directories.
