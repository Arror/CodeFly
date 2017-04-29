all:

init:
	go get -u github.com/samuel/go-thrift/parser
	go get -u github.com/urfave/cli

gen_swift_test:
	./CodeFly json -l swift -i ./sample/sample.thrift -o ./sample/swift

test: clean test_build gen_swift_test

buildTpl:
	rm -rf ./templates/templates.go
	go-bindata -pkg templates -o ./templates/templates.go templates/*/*.tpl

test_build: clean buildTpl
	go build

build: init clean buildTpl
	go build

clean:
	go clean
	rm -rf CodeFly
	rm -rf ./sample/swift