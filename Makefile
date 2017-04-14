.PHONY: all build vendor test lint
# 3
all: lint test
# 
build: vendor
	gb build
# сначало 1 
vendor:
	gb vendor restore
# потом 2 
test: vendor
	gb test -v

lint:
	golint ./src/...
	# gometalinter --disable-all --enable=errcheck --enable=vet --enable=vetshadow ./src/...