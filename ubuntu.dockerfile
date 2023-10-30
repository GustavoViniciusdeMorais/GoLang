FROM gustavovinicius/golang:latest
# FROM grpccplus:latest
# RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
# RUN apt update
# RUN apt install g++
# RUN apt install build-essential
ENTRYPOINT ["tail", "-f", "/dev/null"]