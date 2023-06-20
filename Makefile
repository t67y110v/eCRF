BINARY_NAME=server.exe
USER_NAME=vova


run:	
	go run cmd/server/main.go


test:
	go test .



build: 
	go build -o ${BINARY_NAME} cmd/server/main.go

swag:	
	/home/${USER_NAME}/go/bin/swag init -g cmd/server/main.go

docker.build:
	sudo docker build -t server  .


dc-build:
	sudo docker-compose up --build server 


dc-run:
	sudo docker-compose up server

pr-up:
	sudo docker run \
    -p 9090:9090 \
    prom/prometheus