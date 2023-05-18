build:	
	@go build -o bin/microservices2

run: build 
	@./bin/microservices2

test: 
	@go test -v ./..

btd:
	@cp ./bin/microservices2 ./docker/





