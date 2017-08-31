DEP=~/go/bin/dep
GO=go
FINAL=/usr/local/bin
VERSION=0.0.1-beta
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

all:
	# Dep is golang's official dependancy management test, the /vendor directory(created by dep ensure) can act as another GOPATH binary directory
	$(DEP) ensure
	$(GO) build $(LDFLAGS)
	@mkdir ~/.timeline

install:
	mv timeline $(FINAL)
clean:
	rm timeline
	rm -r vendor/

uninstall:
	rm $(FINAL)/timeline
