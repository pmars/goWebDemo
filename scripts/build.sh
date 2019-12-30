#!/usr/bin/env bash

#!/usr/bin/env bash

project_path=$(pwd)
app_path=${project_path}"/../.."
app_name=pkgame

gitCommit=$(git rev-parse --short HEAD)
buildDate=$(date "+%Y-%m-%d %H:%M:%S")
backDate=$(date "+%Y%m%d%H%M%S")
tag=$(git describe --abbrev=0 --tags)
branch=$(git rev-parse --abbrev-ref HEAD)
goVersion=$(go version)

GOOS=linux ARCH=amd64 go build  -ldflags "-X 'main.buildDate=${buildDate}' -X 'main.gitCommit=${gitCommit}' -X 'main.tag=${tag}' -X 'main.branch=${branch}' -X 'main.goVersion=${goVersion}'" -o ${app_path}/${app_name} ${project_path}/app/main.go

if [[ $? -ne 0 ]]; then
    echo "Build ERROR, Exit Now..."
    exit 1
fi

echo "build success, file path:"${app_path}/${app_name}

echo "exist now..."