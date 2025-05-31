FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod init go-business-card && go mod tidy

CMD ["go", "run", "main.go"]
