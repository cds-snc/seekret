FROM partlab/ubuntu-golang AS build-env
ENV GOOS=linux

RUN apt-get update
RUN apt-get install -y build-essential pkg-config libgit2-dev
RUN mkdir /app
ENV SEEKRET_RULES_PATH=/app/rules
COPY . /app/
WORKDIR /app
