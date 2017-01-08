test:
	docker run --rm -v `pwd`/web:/go/src/web -w /go/src/web golang:latest sh -c 'go get -t && go test -v'

build_dev:
	docker-compose build

run_dev:
	docker-compose up web

clear:
	rm -rf ./web/main

build_prod:
	rm -rf ./web/main
	docker run --rm -v `pwd`/web:/web -w /web golang:latest sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o main .'
