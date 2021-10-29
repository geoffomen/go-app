# note: call scripts from /scripts

help:
	@echo "run 'make myapp' to build"
	@echo "run 'make clean' to clean"

all: myapp

myapp:
	@sh scripts/build.sh myapp

example:
	@sh scripts/build.sh example

clean:
	@sh scripts/clean.sh