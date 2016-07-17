all:

init:

codefly-gen-test:
	./Codefly g -l swift -i inputPath/input.thrift -o outputPath/outputPath

codefly-resize-test:
	./Codefly r -i /Users/Arror/Desktop/abcd.png

build:
	go build Codefly
    
clean:
	go clean
	rm -rf Codefly
	rm -rf *.png