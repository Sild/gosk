#!/bin/bash -e
last_git_tag="$(git tag -l | head -n 1)"
if [ "$1" == "" ]; then
    echo "Usage: $0 vX.Y.Z // to release vX.Y.Z"
    echo "Latest version: ${last_git_tag}"
    exit 1
fi

new_git_tag="$1"

read -p "going to publish version: ${new_git_tag}. Current version: ${last_git_tag} (y/n): " response
if [ "${response}" != "y" ]; then
    echo "Interrupter."
    exit 0
fi
go mod tidy
go build ./...
go test ./...
git commit -a -m "go mod tidy" && git push || true

git tag ${new_git_tag}
git push origin ${new_git_tag}
GOPROXY=proxy.golang.org go list -m github.com/sild/gosk@${new_git_tag}
