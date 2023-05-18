FROM golang:1.16-alpine 


WORKDIR /app



COPY go.mod go.sum ./
COPY .env ./
RUN go mod download 

COPY ./src ./
COPY main.go ./

RUN set CGO_ENABLED=0
RUN set GOOS=linux
RUN go build -o /microservices2



EXPOSE 9000

CMD [ "/microservices2" ]









#syntax=docker/dockerfile:1
# FROM golang:1.16-alpine
# WORKDIR /app
# COPY microservices2/src/ .
# RUN go get -d -v golang.org/x/net/html \
#   && CGO_ENABLED=0 go build -a -installsuffix cgo -o main .





# syntax=docker/dockerfile:1
# FROM golang:1.16-alpine as builder 

# WORKDIR /app

# COPY microservices2 /app/microservices2 



# FROM alpine:latest AS production 

# COPY --from=builder /app /app


# EXPOSE 9000 

# CMD [ "/app/microservices2" ]