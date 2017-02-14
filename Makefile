all:

init:

gen_swift_test:
	./Codefly json -l swift -i ./sample/Base.thrift -o ./sample/swift

test: clean build gen_swift_test

buildTpl:
	rm -rf ./templates/templates.go
	go-bindata -pkg templates -o ./templates/templates.go templates/...

build: clean buildTpl
	go build Codefly

clean:
	go clean
	rm -rf Codefly
	rm -rf ./sample/swift