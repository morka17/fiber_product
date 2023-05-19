build:	
	@go build -o bin/microservices2

run: build 
	@./bin/microservices2

recompile: 
	CGO_ENABLED=0 GOOS=linux go build -o bin/fiber_product
    

test: 
	@go test -v ./..

btd:
	@cp ./bin/microservices2 ./docker/





