test:
	docker-compose run --rm web sh -c 'go get -t && go test -v'

build:
	docker-compose stop web
	docker-compose up -d web
	docker-compose exec web go install
	docker-compose restart web

get:
	docker-compose exec web go get
