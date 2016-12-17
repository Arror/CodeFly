all:

init:

buildTpl:
	rm -rf ./lang/swift/swift_tpl.go
	go-bindata -pkg swift -o ./lang/swift/swift_tpl.go lang/swift/tpl

build:
	go build Codefly

gen:
	./Codefly json -l swift -i /Users/Arror/thrift/Base.thrift -o ./outputPath

help:
	./Codefly -h
    
clean:
	go clean
	rm -rf Codefly
	rm -rf outputPath