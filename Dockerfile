FROM golang:1.19.1-alpine as local

WORKDIR /app

# install golang migrate (migration tool)
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.0

# install air (hot reload tool)
RUN go install github.com/cosmtrek/air@v1.40.2

COPY . .

COPY go.mod go.sum ./

RUN go mod download

CMD [ "air" ]