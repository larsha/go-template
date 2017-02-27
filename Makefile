.PHONY: run redis locust scratch

run:
	REDIS_URI=localhost:6379 PORT=3000 go run main.go

redis:
	docker run --rm -d \
		--name redis \
		-p 6379:6379 \
		redis:latest

locust:
	docker run --rm \
		-v `pwd`/devops/locust:/locust \
		-p 8089:8089 \
		fredriklack/docker-locust \
		locust -H http://localhost:3000

scratch:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .build/main .

	curl -o .build/ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt
	docker build -t go-template .
	rm -rf .build

	docker run --rm -it \
		-p 3000:3000 \
		go-template
