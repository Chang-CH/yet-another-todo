FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy .env file
COPY . . 
COPY .env . 

RUN go build -o main .

CMD ["./main"]
