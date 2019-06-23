prefix ?= /usr/local

.PHONY: all install uninstall

all:
	binpkg download

install:
	binpkg install -p $(prefix)

uninstall:
	binpkg uninstall -p $(prefix)
