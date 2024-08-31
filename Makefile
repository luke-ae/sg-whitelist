build:
	go build -o build/server ./cmd/*.go 

gengql:
	go run github.com/Yamashou/gqlgenc generate .

run:
	go run cmd/* .