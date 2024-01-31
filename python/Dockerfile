FROM python:3.11-slim-buster

ENV PYTHONUNBUFFERED 1
ENV PYTHONDONTWRITEBYTECODE 1
ARG ENV

RUN mkdir -p /python/

# Security updates & Update pip and setuptools
RUN apt-get -y update \
    && apt-get -y upgrade \
    && python -m pip install --upgrade pip setuptools \
    && rm -rf /var/lib/apt/lists/*

# Install python requirements
COPY requirements.txt .
RUN python -m pip install -r requirements.txt

COPY pyproject.toml /python/pyproject.toml
COPY setup.cfg /python/setup.cfg

COPY python/ /python/
WORKDIR /python/

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.11.0/wait /wait
RUN chmod +x /wait
