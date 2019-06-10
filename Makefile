prefix ?= /usr/local
exec_prefix ?= $(prefix)
bindir ?= $(exec_prefix)/bin

.PHONY: all install uninstall fmt test update-vendor

all:
	env GO111MODULE=on go build -mod vendor -v .

install:
	env GO111MODULE=on GOBIN=$(bindir) go install -mod vendor -v .

uninstall:
	rm -f $(bindir)/binpkg

fmt:
	mdfmt -i -w README.md

test:
	go get github.com/frankbraun/gocheck
	gocheck -g -c

update-vendor:
	rm -rf vendor
	env GO111MODULE=on go get -u
	env GO111MODULE=on go mod tidy -v
	env GO111MODULE=on go mod vendor
