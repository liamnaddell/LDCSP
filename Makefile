DEP=$(GOPATH)/bin/dep
GO=go
FINAL=/usr/local/bin
VERSION=0.0.1
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"
PROGNAME=ldcsp
all:
	# Dep is golang's official dependancy management tool, the /vendor direcotry(created by dep ensure) can act as another GOPATH source directory
	$(DEP) ensure
	cd cmd/$(PROGNAME); $(GO) build $(LDFLAGS)
	mv cmd/$(PROGNAME)/$(PROGNAME) .
	

install:
	cp $(PROGNAME) $(FINAL)
	sudo chmod a+x $(FINAL)/$(PROGNAME)

clean:
	rm -f $(PROGNAME)
	rm -rf vendor/

uninstall:
	rm $(FINAL)/$(PROGNAME)
