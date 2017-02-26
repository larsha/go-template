package = github.com/larsha/go-template

test:
	docker-compose run --rm web sh -c 'cd src/${package} && go get -t && go test -v ./...'

# Usage make lib=<lib> get
get:
	docker-compose run --rm web go get ${lib}

format:
	docker-compose run --rm web go fmt ${package}
