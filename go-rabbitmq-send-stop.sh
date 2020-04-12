#!/usr/bin/env bash

# Mysql
echo "Finalizando o mysql..."
docker rm -f mysqldb

# RabbitMQ
echo "Finalizando o rabbitmq..."
docker rm -f rabbitmq

# Teste de conexao
#echo "Finalizando o go-rabbitmq-send..."
#docker rm -f marceloagmelo/go-rabbitmq-send

# Remover rede
echo "Removendo a rede rabbitmq-net..."
docker network rm rabbitmq-net
