imagename := ypjama/niilopaikka

build:
	mkdir -p bin
	go build -o bin/app main.go

build-upx:
	mkdir -p bin
	go build -o bin/app -ldflags='-s' main.go
	upx -9 bin/app

build-docker:
	docker build . -t $(imagename):$$(git rev-parse --short HEAD)
	docker tag $(imagename):$$(git rev-parse --short HEAD) $(imagename):latest

run:
	go run main.go

test:
	go fmt ./...
	go vet ./...
	go test ./...
