FROM golang:1.25

WORKDIR /app

COPY ./backend/services/ ./
RUN go mod download

CMD ["go", "run", "./cmd/main.go"]