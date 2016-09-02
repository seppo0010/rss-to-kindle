FROM golang:alpine

RUN mkdir /go/src/rss-to-kindle
ADD . /go/src/rss-to-kindle
WORKDIR /go/src/rss-to-kindle

RUN apk update
RUN apk add wget git

# download kindlegen and install it to /usr/bin
RUN wget http://kindlegen.s3.amazonaws.com/kindlegen_linux_2.6_i386_v2_9.tar.gz -O /tmp/kindlegen_linux_2.6_i386_v2_9.tar.gz
RUN tar -xzf /tmp/kindlegen_linux_2.6_i386_v2_9.tar.gz -C /tmp
RUN mv /tmp/kindlegen /usr/bin
RUN rm -r /tmp/*

RUN go get github.com/fatih/color
RUN go get github.com/joho/godotenv
RUN go get github.com/mmcdole/gofeed
RUN go get github.com/nfnt/resize
RUN go get github.com/scorredoira/email

RUN go build -o main .
CMD ["/go/src/rss-to-kindle/main"]