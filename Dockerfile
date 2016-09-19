FROM golang:latest
MAINTAINER https://github.com/jfmyers9/cf-summit-sample-app

RUN mkdir /sample-app
ADD . /sample-app/
WORKDIR /sample-app

RUN go build -o cf-summit .

EXPOSE 8080

CMD ["/sample-app/cf-summit"]
