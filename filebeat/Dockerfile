FROM alpine:3.8
LABEL maintainer="will835559313@163.com"
RUN apk update \
	&& apk add ca-certificates wget \
	&& update-ca-certificates \
	&& wget https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-6.3.2-linux-x86_64.tar.gz \
	&& tar xf filebeat-6.3.2-linux-x86_64.tar.gz \
	&& mv filebeat-6.3.2-linux-x86_64 /usr/local/filebeat \
    && addgroup filebeat \
	&& adduser filebeat -G filebeat \
    && mv filebeat.yml filebeat.yml.ori \
    && wget https://raw.githubusercontent.com/mgxian/k8s-log/master/filebeat/filebeat.yml \
    && chown filebeat.filebeat -R /usr/local/filebeat \
	&& apk del build-dependencies \
	&& rm -rf /var/cache/apk/*
WORKDIR /usr/local/filebeat
USER filebeat
ENTRYPOINT [ "./filebeat" ]
CMD [ "-e", "-c", "./filebeat.yml" ]