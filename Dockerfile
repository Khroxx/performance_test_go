FROM golang:1.25

WORKDIR /app
COPY . .

RUN go mod download
RUN go install github.com/air-verse/air@latest

EXPOSE 8081

CMD ["air"]