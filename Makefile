test:
	docker-compose run --rm web sh -c 'go get -t && go test -v'

build:
	docker-compose stop web
	docker-compose up -d web
	docker-compose exec web go install
	docker-compose restart web

get:
	docker-compose exec web go get

# prod:
# 	docker run --rm -v `pwd`/web:/web -w /web golang:latest sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o app .'
