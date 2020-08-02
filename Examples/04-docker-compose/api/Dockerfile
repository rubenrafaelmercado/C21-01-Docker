FROM python:3.8.4-alpine

WORKDIR /app

ADD requirements.txt /app

RUN pip install -r requirements.txt

EXPOSE 5000

COPY src/ /app/

ENTRYPOINT ["/usr/local/bin/python", "app.py"]
