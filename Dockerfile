FROM golang:1.12.0-stretch AS build-env
ENV GOOS=linux

RUN apt-get update
RUN apt-get install -y build-essential=12.3 --no-install-recommends
RUN mkdir /app

WORKDIR /app

RUN go get github.com/cds-snc/seekret
RUN go build -a -installsuffix cgo -o seekret github.com/cds-snc/seekret
RUN chmod +x /app/seekret

FROM debian:stretch-slim
ENV SEEKRET_RULES_PATH /rules
RUN mkdir /rules
COPY ./rules /rules
COPY --from=build-env /app/seekret /bin/
CMD ["/bin/seekret"]
