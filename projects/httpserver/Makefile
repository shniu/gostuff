export tag=v1.0.0

root:
	export ROOT=github.com/shniu/cncamp-golang

build:
	echo "building httpserver binary for macos ..."
	mkdir -p bin/macos/amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/macos/amd64

build_linux:
	echo "building httpserver binary for linux ..."
	mkdir -p bin/linux/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64

release: build_linux
	echo "building httpserver container ..."
	docker build -t cncamp/httpserver:${tag} .

push: release
	echo "pushing cncamp/httpserver"
	docker push cncamp httpserver:${tag}

