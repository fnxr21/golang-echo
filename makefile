.PHONY: run db
run:
	go run main.go

# Start the Docker containers in the 'db' directory
dc:
	docker-compose up 
dc-up:
	docker-compose up -d
# make dc-build TAG=1.0.0
dc-build:
	docker build -t nxrfandi/nexmedis:${TAG} .

dc-push:
	docker push nxrfandi/nexmedis:${TAG}

test:
	go test -v ./internal/handler

