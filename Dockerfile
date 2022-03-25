FROM ubuntu:latest

RUN mkdir /certs
RUN mkdir /config

EXPOSE 443/tcp
EXPOSE 8080/tcp

ADD main /main

CMD /main
