all:

init:

buildTpl:
	rm -rf ./lang/swift/swift_tpl.go
	go-bindata -pkg templates -o ./templates/templates.go templates/swift

build:
	go build Codefly

gen:
	./Codefly json -l swift -i /Users/Arror/thrift/Base.thrift -o ./outputPath

test: clean buildTpl build gen

help:
	./Codefly -h
    
clean:
	go clean
	rm -rf Codefly
	rm -rf outputPath