FROM ubuntu:jammy
# RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
RUN apt update
RUN apt install curl -y
RUN curl -OL https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
RUN sha256sum go1.16.7.linux-amd64.tar.gz
RUN tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin
RUN go version
RUN cd / && mkdir go