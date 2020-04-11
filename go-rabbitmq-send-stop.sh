#!/usr/bin/env bash

# Mysql
echo "Finalizando o mysql..."
docker rm -f mysqldb

# RabbitMQ
echo "Finalizando o rabbitmq..."
docker rm -f rabbitmq

# Teste de conexao
#echo "Finalizando o go-teste-conexao..."
#docker rm -f marceloagmelo/go-teste-conexao

# Remover rede
echo "Removendo a rede teste-conexao-net..."
docker network rm rabbitmq-net
