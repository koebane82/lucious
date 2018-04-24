FROM "centos:centos6"

MAINTAINER koebane

RUN yum install -y gcc \
&& curl -o /tmp/golang.tgz https://dl.google.com/go/go1.10.1.linux-amd64.tar.gz \
&& cd /usr/local \
&& tar xzf /tmp/golang.tgz \
&& mkdir -p /usr/local/repo

ENV PATH="/usr/local/go/bin:${PATH}"

COPY files/gobuild.sh  /usr/local/gobuild.sh
RUN chmod 755 /usr/local/gobuild.sh

ENTRYPOINT ["/usr/local/gobuild.sh"] 