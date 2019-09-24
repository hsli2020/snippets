######## build.sh

#!/usr/bin/env bash

mkdir -p 'builds'

cd 'cmd/upload'

go get -d -v ./...
go install -v ./...

package_name=upload

platforms=("windows/amd64" "linux/amd64" "darwin/amd64" "linux/arm" )

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH

    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH GOARM=7 go build -o $output_name .

    sha256sum $output_name >> SHA256SUMS

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi

    mv $output_name ../../builds
done

mv SHA256SUMS ../../builds

######## upload.sh

#!/bin/bash

HOST=http://127.0.0.1:9429
TOKEN=YOUR_TOKEN

if [ $# -eq 0 ]
  then
    echo 'No argument supplied'
    exit 1
  else
    url=$(curl --silent --header "Token: $TOKEN" -F ''file=@"$1"'' $HOST)
    echo $url
fi
