FROM golang
WORKDIR /go/src/musicstore-go-api
ADD . /go/src/musicstore-go-api
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go build main.go
EXPOSE 3000
ENTRYPOINT ./main