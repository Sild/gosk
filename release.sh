#!/bin/bash -e
if [ "$1" == "" ]; then
    echo "Usage: $0 x.y.z // to release vX.Y.Z"
    exit 1
fi

git_tag="v$1"
read -p "going to publish version: ${git_tag} (y/n): " response
if [ "${response}" != "y" ]; then
    echo "Interrupter."
    exit 0
fi
go mod tidy
git commit -a -m "go mod tidy" || true
go build ./...
go test ./...

git tag v0.1.0
git push origin v0.1.0
