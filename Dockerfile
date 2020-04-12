FROM golang:1.13.6

ENV APP_HOME /go/src/github.com/marceloagmelo/go-rabbitmq-send

ADD . $APP_HOME

WORKDIR $APP_HOME

RUN go get ./... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o go-rabbitmq-send && \
    chown -fR $USER:0 $APP_HOME && \
    rm -Rf /tmp/* && rm -Rf /var/tmp/*

#######################################################################
##### We have to expose image metada as label and ENV
#######################################################################
LABEL br.com.santander.imageowner="Corporate Techonology" \
      br.com.santander.description="RabbitMQ envio de mensagens runtime for node microservices" \
      br.com.santander.components="Golang Server"

ENV br.com.santander.imageowner="Corporate Techonology"
ENV br.com.santander.description="RabbitMQ envio de mensagens runtime for node microservices"
ENV br.com.santander.components="Golang Server"

ENV PATH $APP_HOME:$PATH

EXPOSE 8080

CMD [ "go-rabbitmq-send" ]