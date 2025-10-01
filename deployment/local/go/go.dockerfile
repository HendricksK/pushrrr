FROM golang:latest 

RUN export GOPATH=$HOME/go

RUN export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

RUN apt update -y && apt install supervisor -y

RUN mkdir -p /var/log/supervisor

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY /app/. .

RUN go mod tidy

COPY /proc/. /etc/supervisor/conf.d/

EXPOSE 9000


CMD ["air", "-c", "dev.air.toml"]
# CMD ["/usr/bin/supervisord"]

# CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/supervisord.conf"]

