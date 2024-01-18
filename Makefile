install:
	go mod install

generate:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

build:
	go build main.go

run:
	go run main.go server

docker-build:
	docker build -t geoproperty-be .

docker-run:
	docker run -d --name -p 3000:3000 --env-file .env geoproperty-be