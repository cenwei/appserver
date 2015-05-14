Welcome to Sharelog AppServer 

[![Circle CI](https://circleci.com/gh/sharelog/appserver.svg?style=svg&circle-token=a52823aea372317d16270f0eade6f9a8d8bb1ca9)](https://circleci.com/gh/sharelog/appserver)

Dependencies
============
1. go v1.4
2. https://github.com/tools/godep is used for managing go lib

Development
===========
$ `godep restore && go build && ./appserver config.ini`

config.ini should be provided in args.

Test
====
You may refer to .travis.yml

$ `godep restore && go test github.com/sharelog/appserver/...`

For local development, you are suggested to open GoConvey to keep track of testing status.

$ `go get github.com/smartystreets/goconvey`
