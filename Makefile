BINARY ?= docker-rmi
BINDIR ?= $(DESTDIR)/usr/local/bin
GOLANG ?= /usr/local/go/bin/go

.PHONY: build
build: main.go
	$(GOLANG) build -o $(BINARY) .

.PHONY: install
install:
	$(GOLANG) build -o $(BINARY) .
	install -m 755 $(BINARY) $(BINDIR)/$(BINARY)

.PHONY: clean
clean: 
	rm -f $(BINARY)
