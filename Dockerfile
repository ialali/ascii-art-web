FROM golang:1.20-alpine3.18

WORKDIR /app

COPY . .

RUN go build -o ascii-art-web main.go

EXPOSE 5050

CMD [ "/app/main" ]