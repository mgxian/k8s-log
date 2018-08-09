FROM alpine:3.8
LABEL maintainer="will835559313@163.com"
RUN apk add --no-cache wget libc6-compat \
    && wget https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-6.3.2-linux-x86_64.tar.gz \
    && tar xf filebeat-6.3.2-linux-x86_64.tar.gz \
    && mv filebeat-6.3.2-linux-x86_64 /usr/local/filebeat \
    && addgroup filebeat \
    && adduser -D filebeat -G filebeat \
    && chown filebeat.filebeat -R /usr/local/filebeat \
    && rm -rf /var/cache/apk/* filebeat-6.3.2-linux-x86_64.tar.gz
WORKDIR /usr/local/filebeat
USER filebeat
COPY filebeat.yml config/filebeat.yml
ENTRYPOINT [ "./filebeat" ]
CMD [ "-e", "-c", "./config/filebeat.yml" ]
