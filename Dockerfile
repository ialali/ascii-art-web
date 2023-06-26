FROM golang:1.20-alpine3.18

LABEL maintainer="Ibrahim Al-Ali <ibrahim.alali94@yahoo.ro>"
LABEL description="Docker image for the ASCII Art Web application"

WORKDIR /app

COPY . .

RUN go build -o ascii-art-web main.go

EXPOSE 5050

CMD [ "/app/ascii-art-web" ]