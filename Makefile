all:

init:

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