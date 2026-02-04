all: test

GIT_TAG := $(shell git describe --tags --abbrev=0)

test:
	go test -cover -v ./...

publish:
	git push
	git push --tags
	curl https://proxy.golang.org/github.com/fatindeed/tencentcloud-sdk-go/@v/$(GIT_TAG).info

unpublish:
	@echo -n "Are you sure unpublish $(GIT_TAG)? [y/N] " && read ans && [ $${ans:-N} = y ]
	git tag -d $(GIT_TAG)
	git push --delete origin $(GIT_TAG)