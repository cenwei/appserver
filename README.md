Welcome to Sharelog AppServer 

# TODO: add travis token
[![Build Status](https://magnum.travis-ci.com/sharelog/appserver.svg?token=$TOKEN)](https://magnum.travis-ci.com/sharelog/appserver)

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

run `godep restore && go test github.com/sharelog/appserver/...`

For local development, you are suggested to open GoConvey to keep track of testing status.

refs: https://github.com/smartystreets/goconvey
