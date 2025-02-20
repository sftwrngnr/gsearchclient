export GIT_COMMIT:=$(shell git rev-list -1 HEAD)
build:
	@echo $(GIT_COMMIT)
	rm -f bin/gsearch
	go build -ldflags "-X main.GitCommit=${GIT_COMMIT}" -o bin/gsearch ./main.go

run:
	go run ./main.go $(ARGS)

nocleanbuild:
	go build -ldflags "-X main.GitCommit=${GIT_COMMIT}" -o bin/gsearch ./main.go