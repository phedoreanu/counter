# counter [![Build Status](https://travis-ci.org/phedoreanu/counter.svg?branch=master)](https://travis-ci.org/phedoreanu/counter) [![Coverage Status](https://coveralls.io/repos/github/phedoreanu/counter/badge.svg)](https://coveralls.io/github/phedoreanu/counter) [![Go Report Card](https://goreportcard.com/badge/github.com/phedoreanu/counter)](https://goreportcard.com/report/github.com/phedoreanu/counter)

[![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org) [![GoDoc](https://godoc.org/github.com/phedoreanu/counter?status.svg)](https://godoc.org/github.com/phedoreanu/counter)

###Installation
Counter requires Go 1.5 or later.
```go
$ go get -u github.com/phedoreanu/counter
```
###Usage
```go
$ docker-compose up
```

###Unit tests
```go
$ go test -v -race -cover -parallel 8 -cpu 8
```

###Smoke tests
```shell
$ ansible-playbook smoke-tests.yml
```

###Load tests
Start the app and execute:
```go
$ HOSTNAME=localhost:8080 ./load-tests.sh
```
Open `plot.html` for a nice graph.

###Cyclomatic complexity
```go
$ mccabe-cyclomatic -p github.com/phedoreanu/counter
4
```

```go
$ gocyclo -top 3 -avg .
4 main (*Env).SyncCounter main.go:39:1
3 db (*DB).IncrementCounter db/db.go:64:1
3 db (*DB).ReadCounter db/db.go:48:1
Average: 2
```
