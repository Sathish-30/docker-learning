FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main.exe main.go

EXPOSE 8000

CMD [ "./main.exe" ]
