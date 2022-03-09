FROM golang:1.17
WORKDIR /ports 
COPY ./ /ports

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN export PATH=$PATH:$GOPATH/bin

RUN make tidy

CMD go run cmd/main.go