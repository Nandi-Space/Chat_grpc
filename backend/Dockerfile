FROM golang:alpine

WORKDIR /backend
COPY . .
COPY go.mod ./

RUN go mod tidy

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait


ENTRYPOINT export SERVER_ADDRESS=db && CompileDaemon --build="go build cmd/api/main.go" --command="./main"