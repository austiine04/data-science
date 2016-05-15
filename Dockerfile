FROM dataquestio/python3-starter
MAINTAINER Augustine Kisitu austiine04@gmail.com

RUN mkdir -p /tmp/project
COPY .  /tmp/project
WORKDIR /tmp/project

EXPOSE 8888
