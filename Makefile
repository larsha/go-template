test:
	docker-compose run --rm web sh -c 'go get -t && go test -v'

build:
	docker-compose stop web
	docker-compose up -d web
	docker-compose exec web go install
	docker-compose restart web

get:
	docker-compose exec web go get

# Usage: make host=http://localhost locust
locust:
	docker build --build-arg PATH_TO_LOCUST=/devops/locust -t locust -f Dockerfile.locust .
	docker run -p 8089:8089 -it --rm locust /usr/bin/locust -H $(host)
	docker rmi locust
