build:
	rm -f bin/gsearch
	go build -o bin/gsearch ./main.go

run:
	go run ./main.go $(ARGS)

nocleanbuild:
	go build -o bin/gsearch ./main.go