setenv:
	@echo export GOPATH=`pwd`
	@echo export GOBIN=`pwd`/bin

build:
	go get

run: setenv build
	go run main.go
