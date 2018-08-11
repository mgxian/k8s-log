FROM python:3-alpine
LABEL maintainer="will835559313@163.com"

RUN pip install requests
COPY app.py /
ENTRYPOINT [ "python", "/app.py" ]