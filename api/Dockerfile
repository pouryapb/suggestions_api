FROM registry.gitlab.com/qio/standard/golang:1.17

ENV GIN_MODE=release

WORKDIR /api

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build

EXPOSE 80
CMD [ "./suggestions.api" ]
