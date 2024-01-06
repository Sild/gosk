#!/bin/bash -e

last_git_tag="$(git tag -l | head -n 1)"
new_git_tag="$1"

while [ "${new_git_tag}" == "" ]; do 
    echo "Latest version: ${last_git_tag}"
    read -p "Enter next version: " new_git_tag
done


read -p "Going to publish version: ${new_git_tag}. Current version: ${last_git_tag}. Continue? (y/n): " response

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
