FROM python:3.11.6-slim-bookworm

ENV PYTHONUNBUFFERED 1
ENV PYTHONDONTWRITEBYTECODE 1

RUN mkdir -p /src/py_app/

# Security updates & Update pip and setuptools
RUN apt-get -y update \
    && apt-get -y upgrade \
    && python -m pip install --upgrade pip setuptools \
    && rm -rf /var/lib/apt/lists/*

# Install python requirements
COPY requirements.txt .
RUN python -m pip install -r requirements.txt

COPY /src/py_app/ /src/py_app/
WORKDIR /src/py_app/

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.11.0/wait /wait
RUN chmod +x /wait
