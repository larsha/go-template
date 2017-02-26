test:
	docker-compose run --rm web sh -c 'go get -t && go test -v'

run:
	docker-compose stop web
	docker-compose build web
	docker-compose up web
