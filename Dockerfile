From busybox:ubuntu-14.04

RUN apt-get update && apt-get install -y automake

ADD . /mesos-go

RUN cd /mesos-go/examples && make

WORKDIR /mesos-go/examples/_output