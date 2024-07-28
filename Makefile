GOCMD=go
GOFILES=$(shell find . -type f -name "*.go" -not -path "./vendor/*")

GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

APP_NAME=main
BINARY_WIN=.exe

BUILDDIR=s2hWS

RELEASE_NAME=s2hWS
RELEASEDIR=$(BUILDDIR)
RELEASEFILE=$(RELEASE_NAME).zip

all: release

setup:
	mkdir -p $(RELEASEDIR)

clean:
	$(GOCLEAN)
	rm -rf $(RELEASE_NAME) $(BUILDDIR)
	rm -rf .s2hWS

format:
	goimports -w -d $(GOFILES)

test:
	go test -cover ./...

build-win: format clean setup
	rsrc -manifest main.exe.manifest -ico assets/icon.ico -o rsrc.syso
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(APP_NAME).go -o $(RELEASEDIR)/$(RELEASE_NAME)$(BINARY_WIN)

release: build-win
