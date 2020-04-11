#!/usr/bin/env bash

source setenv.sh

# Criar rede
echo "Criando a rede rabbitmq-net..."
docker network create rabbitmq-net 

# Mysql
echo "Subindo o mysql..."
docker run -d --name mysqldb --network rabbitmq-net  \
-p 3306:3306 \
-e MYSQL_USER=${MYSQL_USER} \
-e MYSQL_PASSWORD=${MYSQL_PASSWORD} \
-e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
-e MYSQL_DATABASE=${MYSQL_DATABASE} \
mysql:5.7

# RabbitMQ
echo "Subindo o rabbitmq..."
docker run -d --name rabbitmq --network rabbitmq-net  \
-p 5672:5672 -p 15672:15672 \
-e RABBITMQ_DEFAULT_USER=${RABBITMQ_DEFAULT_USER} \
-e RABBITMQ_DEFAULT_PASS=${RABBITMQ_DEFAULT_PASS} \
-e RABBITMQ_ERLANG_COOKIE=${RABBITMQ_ERLANG_COOKIE} \
-e RABBITMQ_DEFAULT_VHOST=${RABBITMQ_DEFAULT_VHOST} \
rabbitmq:3.6.16-management

# Teste de conexao
#echo "Subindo o go-teste-conexao..."
#docker run -d --name go-teste-conexao --network rabbitmq-net  \
#-p 8080:8080 \
#go-teste-conexao

# Listando os containers
docker ps

# centos/mariadb-102-centos7