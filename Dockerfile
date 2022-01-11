FROM golang:1.17

RUN mkdir /go/src/golang_todo_app
WORKDIR /go/src/golang_todo_app
ADD . /go/src/golang_todo_app

RUN apt-get update && apt-get install -y sqlite3
