test:
	docker-compose run --rm web sh -c 'go get -t && go test -v'

# Usage make package=<package> get
get:
	docker-compose run --rm web go get ${package}

run:
	docker-compose stop web
	docker-compose build web
	docker-compose up web
