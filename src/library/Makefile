.PHONY: help dynamic static main run clean
help:
	@echo "dynamic - build dynamic lib"
	@echo " static - build static lib"

dynamic:
	# compile the source into an object file (position independent code)
	gcc -fPIC -c mylib.c
	# convert the resulting object file into a shared library
	gcc -shared -o libmylib.so mylib.o
# dynamic OMIT

static:
	# compile the source into an object file
	gcc -c mylib.c
	# convert the resulting object file into a library
	ar rc libmylib.a mylib.o
	# build an index inside the library
	ranlib libmylib.a
# static OMIT

main:
	go build main.go

run:
	go run main.go

clean:
	rm -f libmylib.so
	rm -f libmylib.a
	rm -f mylib.o
	rm -f main
