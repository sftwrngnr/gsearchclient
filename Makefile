build:
	go build -o bin/gsearch ./main.go

run:
	go run ./main.go $(ARGS)
