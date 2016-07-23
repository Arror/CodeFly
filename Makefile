all:

init:

codefly-gen-test:
	./Codefly g -l swift -i Example/Example.thrift -o outputPath/outputPath

build:
	go build Codefly

help:
	./Codefly -h
    
clean:
	go clean
	rm -rf Codefly
	rm -rf *.png
	rm -rf outputPath