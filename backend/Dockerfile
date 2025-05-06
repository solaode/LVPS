FROM golang:1.22
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o lvps-backend ./src/main.go
CMD ["./lvps-backend"]
