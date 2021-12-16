FROM tiangolo/uvicorn-gunicorn-fastapi:python3.9-slim
MAINTAINER Daniel Travieso <danielgtravieso@gmail.com>

COPY ./requirements.txt /app/requirements.txt

RUN pip install --no-cache-dir --upgrade -r /app/requirements.txt

ENV RPC_SERVER_ADDRESS "localhost:9090"
ENV APP_MODULE "main:app"
COPY . /app
