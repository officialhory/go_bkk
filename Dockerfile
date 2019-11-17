FROM golang:latest

LABEL maintainer="officialhory"

WORKDIR /app
COPY . .

RUN go build -o main .

EXPOSE 80

CMD ["./main"]
