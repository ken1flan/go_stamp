FROM alpine:latest

MAINTAINER ken1flan <ken1flan@gmail.com>

WORKDIR "/opt"

ADD .docker_build/go_stamp /opt/bin/go_stamp
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/go_stamp"]
