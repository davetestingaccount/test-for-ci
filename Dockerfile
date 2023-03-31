FROM golang:latest

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

# add golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin/

EXPOSE 80

COPY release/api-binary /bin/

ENTRYPOINT ["/bin/api-binary"]