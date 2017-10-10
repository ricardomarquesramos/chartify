FROM alpine:latest

MAINTAINER Ricardo Ramos

WORKDIR "/opt"

ADD .docker_build/chartify /opt/bin/chartify

CMD ["/opt/bin/chartify"]