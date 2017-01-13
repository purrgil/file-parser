#!/usr/bin/env bash

echo "mode: set" > full_coverage.out
for pkg in $(go list ./... | grep -v /vendor/); do
    go test -v -coverprofile=coverage.out -covermode=set $pkg
	if [ $? -ne 0 ]; then
		exit 1
	fi
    grep -h -v "^mode: set" coverage.out >> full_coverage.out
done

#go tool cover -html=full_coverage.out -o=coverage.html
go tool cover -func full_coverage.out
rm -f coverage.out 

exit 0