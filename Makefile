package = github.com/larsha/go-template

test:
	docker-compose run --rm web sh -c 'cd src/${package} && go get -t && go test -v ./...'

# Usage make lib=<lib> get
get:
	docker-compose run --rm web go get ${lib}

format:
	docker-compose run --rm web go fmt ${package}

scratch:
	docker run --rm \
		-v `pwd`:/go/src/main \
		-w /go/src/main \
		golang:latest \
		sh -c 'go get && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .build/main .'

	curl -o .build/ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt
	docker build -t go-template -f Dockerfile.scratch .
	rm -rf .build
	docker run -it -p 3000:3000 go-template
