FROM golang:1.11

RUN groupadd -r simplewebapp && useradd -r -g simplewebapp simplewebapp
RUN mkdir /go/src/simple-web-app
WORKDIR /go/src/simple-web-app
COPY . .

RUN go get github.com/DavidHuie/gomigrate
RUN go get github.com/lib/pq
RUN go build ./cmd/simple-web-app/main.go

#EXPOSE 1988

USER simplewebapp
CMD ["./main"]