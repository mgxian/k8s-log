FROM openjdk:8-jre-alpine
LABEL maintainer="will835559313@163.com"
RUN apk add --no-cache wget \
    && wget https://artifacts.elastic.co/downloads/logstash/logstash-6.3.2.tar.gz \
    && tar xf logstash-6.3.2.tar.gz \
    && mv logstash-6.3.2 /usr/local/logstash \
    && addgroup logstash \
    && adduser -D logstash -G logstash \
    && chown logstash.logstash -R /usr/local/logstash \
    && rm -rf /var/cache/apk/* logstash-6.3.2-linux-x86_64.tar.gz
WORKDIR /usr/local/logstash
USER logstash
COPY shipper.conf config/logstash.conf
ENTRYPOINT [ "./bin/logstash" ]
CMD [ "-f", "./config/logstash.conf" ]