security-scan:
	gosec ./..

build-server:
	CGO_ENABLED=true go build -o ./bin/cnf-q-service ./cmd/server/main.go

build-cli:
	CGO_ENABLED=true go build -o ./bin/cnf-q-cli ./cmd/cli/main.go

docker-build:
	docker build . --file ./deploy/docker/Dockerfile -t cnfsoftware/cnf-q-service

docker-push:
	docker push cnfsoftware/cnf-q-service