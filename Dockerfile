FROM golang

RUN apt-get -y update
RUN apt-get -y upgrade

WORKDIR /root/study-a-tour-of-go
