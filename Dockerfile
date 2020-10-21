FROM golang:latest

ENV NAME rikiapi
ENV PORT 80

COPY . /go/src/${NAME}
WORKDIR /go/src/${NAME}

RUN go get ./
RUN go build -o ${NAME}

CMD ./${NAME}

EXPOSE ${PORT}