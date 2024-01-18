install:
	pip install -r requirements.txt

generate:
	python -m grpc_tools.protoc -I./proto --python_out=./service --pyi_out=./service --grpc_python_out=./service ./proto/*.proto

run:
	python -W ignore -m jurigged -v app.py

docker-build:
	docker build -t geoproperty-gpt .

docker-run:
	docker run -d --name geoproperty-gpt -p 50051:50051 --env-file .env geoproperty-gpt