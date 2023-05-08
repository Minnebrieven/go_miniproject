FROM golang:1.20-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app