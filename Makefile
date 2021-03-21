all: lib main

lib:
	gcc -o libtest_c.so -fPIC -shared -Wall test.c

main:
	go build main.go