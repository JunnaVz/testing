FROM golang:1.21

WORKDIR /app

COPY . .

CMD ["go", "test", "./..."]

#CMD ["go", "run", "./cmd/main.go"]