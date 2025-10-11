#########################################################################
# File Name: test.sh
# Author: frank
# mail: 1216451203@qq.com
# Created Time: 2025年09月12日 星期五 15时51分36秒
#########################################################################
#!/bin/bash

curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john",
    "email": "john@example.com",
    "password": "secret123"
  }'
echo ""

curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "tom",
    "email": "tom@example.com",
    "password": "secret123"
  }'
echo ""

curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john",
    "password": "secret123"
  }'
echo ""

curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "tom",
    "password": "secret123"
  }'
echo ""

curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "jack",
    "password": "secret123"
  }'
echo ""

