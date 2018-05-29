FROM golang:1.10.2

WORKDIR /go/src/campaigns
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
# RUN go get github.com/codegangsta/gin
