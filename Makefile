all:

init:

codefly-gen-test:
	./Codefly g -l swift -i inputPath/input.thrift -o outputPath/outputPath

build:
	go build Codefly
    
clean:
	go clean
	rm -rf Codefly