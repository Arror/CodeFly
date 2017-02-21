all:

init:

gen_swift_test:
	./CodeFly json -l swift -i ./sample/sample.thrift -o ./sample/swift

test: clean build gen_swift_test

buildTpl:
	rm -rf ./templates/templates.go
	go-bindata -pkg templates -o ./templates/templates.go templates/*/*.tpl

build: clean buildTpl
	go build

clean:
	go clean
	rm -rf CodeFly
	rm -rf ./sample/swift