.PHONY: build
VERSION=0.0.0

build: COMMIT=$(shell git rev-list -1 HEAD | grep -o "^.\{10\}")
build: DATE=$(shell date +'%Y-%m-%d %H:%M')
build: 
	env GOOS=darwin  GOARCH=amd64 go build -mod vendor -ldflags '-X "main.Version=$(VERSION) ($(COMMIT) - $(DATE))"' -o build/$(VERSION)/grok-$(VERSION)-darwin
	env GOOS=linux   GOARCH=amd64 go build -mod vendor -ldflags '-X "main.Version=$(VERSION) ($(COMMIT) - $(DATE))"' -o build/$(VERSION)/grok-$(VERSION)-linux
	env GOOS=windows GOARCH=amd64 go build -mod vendor -ldflags '-X "main.Version=$(VERSION) ($(COMMIT) - $(DATE))"' -o build/$(VERSION)/grok-$(VERSION)-windows.exe

publish:
	mkdir -p /keybase/public/justmiles/artifacts/grok/$(VERSION)/
	cp build/$(VERSION)/grok-$(VERSION)-linux /keybase/public/justmiles/artifacts/grok/$(VERSION)/grok-linux