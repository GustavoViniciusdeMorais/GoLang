FROM gustavovinicius/golang:nginx
# FROM grpccplus:latest
# RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
# RUN apt update
# RUN apt install g++
# RUN apt install build-essential
# RUN export PATH=$PATH:/usr/local/go/bin
WORKDIR /var/www/html
ENTRYPOINT ["tail", "-f", "/dev/null"]