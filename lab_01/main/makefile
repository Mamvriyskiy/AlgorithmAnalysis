.PHONY: all c go run 

TESTLIBPATH="./ctestlib"

all: c go run


c:
	gcc -c -Wall -Werror -fpic -o testtime.o testtime.c
	gcc -shared -o C.so testtime.o

go:
	go build -o app *.go

run:
	./app