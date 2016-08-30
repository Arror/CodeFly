all:

init:

build:
	go build Codefly

help:
	./Codefly -h
    
clean:
	go clean
	rm -rf Codefly
	rm -rf outputPath