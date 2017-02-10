FROM  github.com/jrlmx2/stockAnalysis:latest

RUN apt-get -y update && apt-get -y upgrade

RUN wget https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz

RUN tar -xzf go1.7.linux-amd64.tar.gz
RUN export GOROOT=/go
RUN export GOPATH=
