FROM alpine:latest

MAINTAINER Ricardo Ramos

WORKDIR "/opt"

ADD .docker_build/go-chartify /opt/bin/go-chartify

CMD ["/opt/bin/go-chartify"]