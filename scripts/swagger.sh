#!/bin/sh

set -e

swag fmt

if [ ! -f internal/version/VERSION ]; then
    echo "internal/version/VERSION not exists"
    exit 1
fi

VERSION=$(cat internal/version/VERSION)
echo "Version: $VERSION"

if [ ! -f main.go ]; then
    echo "main.go not exists"
    exit 1
fi

if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' "s|//	@version		.*|//	@version		$VERSION|" main.go
else
    sed -i "s|//	@version		.*|//	@version		$VERSION|" main.go
fi
cat main.go | grep $VERSION > /dev/null || (echo "update version failed" && exit 1)

swag init -o support_files/api_docs --parseDependency --parseInternal
