FROM fluent/fluentd:v1.2-debian
LABEL maintainer="will835559313@163.com"

RUN fluent-gem install fluent-plugin-elasticsearch \
    && mkdir /var/log/fluentd-buffers \
    && chown -R fluent:root /var/log/fluentd-buffers