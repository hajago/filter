# filter

[![Build Status](https://travis-ci.org/travis-ci/travis-web.svg?branch=master)](https://travis-ci.org/travis-ci/travis-web)

## try development

1. Install Go and set GOROOT, path
	1. http://golang.org/doc/install

1. Set GOPATH
	1. http://golang.org/doc/code.html#GOPATH
	
1. Get project

		$ go get github.com/hajago/filter
		$ cd $GOPATH/src/github.com/hajago/filter
	
1. Install Godep 

		$ go get github.com/tools/godep
		
1. Install filter project

		$ godep go install
		$ filter
