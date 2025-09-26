FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/auction cmd/auction/main.go

ENV BATCH_INSERT_INTERVAL="20s"
ENV MAX_BATCH_SIZE="4"
ENV AUCTION_INTERVAL="20s"
ENV MONGODB_URL="mongodb://admin:admin@mongodb:27017/auctions?authSource=admin"
ENV MONGODB_DB="auctions"

EXPOSE 8080

ENTRYPOINT ["/app/auction"]
