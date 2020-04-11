#FROM marceloagmelo/golang-1.13 AS builder
FROM golang:1.13.6 AS builder

USER root

ENV APP_HOME /go/src/github.com/marceloagmelo/go-teste-conexao

COPY Dockerfile $IMAGE_SCRIPTS_HOME/Dockerfile
#COPY .netrc /home/golang/.netrc
ADD . $APP_HOME

WORKDIR $APP_HOME

RUN go mod init && \
    #CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-teste-conexao && \
    go install && \
    rm -Rf /tmp/* && rm -Rf /var/tmp/*

###
# IMAGEM FINAL
###
FROM scratch

ENV GOLANG_VERSION 1.13.6
ENV APP_HOME /opt/app

RUN mkdir -p $APP_HOME
COPY Dockerfile $APP_HOME/Dockerfile

WORKDIR $APP_HOME
COPY --from=builder $APP_HOME/go-teste-conexao .
COPY views $APP_HOME/views
COPY static $APP_HOME/static

#######################################################################
##### We have to expose image metada as label and ENV
#######################################################################
LABEL br.com.santander.imageowner="Corporate Techonology" \
      br.com.santander.description="Teste de conexao runtime for node microservices" \
      br.com.santander.components="Golang Server"

ENV br.com.santander.imageowner="Corporate Techonology"
ENV br.com.santander.description="Teste de conexao runtime for node microservices"
ENV br.com.santander.components="Golang Server"

EXPOSE 8080

CMD [ "go-teste-conexao" ]
