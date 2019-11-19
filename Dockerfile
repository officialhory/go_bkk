FROM golang:latest

LABEL maintainer="officialhory"

WORKDIR /app
COPY . .

RUN go build -o main .

EXPOSE 6000

CMD ["./main"]
