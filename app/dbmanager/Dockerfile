FROM golang:latest

WORKDIR /build 

COPY . .

RUN go mod tidy
RUN go build -o /app/dbmanager dbmanager/main.go

WORKDIR /app
RUN rm -r /build


CMD [ "/app/dbmanager" ]