security-scan:
	gosec ./..

build:
	CGO_ENABLED=true go build -o ./bin/queue-service ./cmd/server/main.go

docker-build:
	docker build . --file ./deploy/docker/Dockerfile -t cnfsoftware/cnf-q-server

docker-push:
	docker push cnfsoftware/cnf-q-server