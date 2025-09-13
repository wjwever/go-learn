#########################################################################
# File Name: test.sh
# Author: frank
# mail: 1216451203@qq.com
# Created Time: 2025年09月12日 星期五 15时51分36秒
#########################################################################
#!/bin/bash
set -x


# 根据登录返回的token进行修改
token1="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTc4NDMxOTIsInVzZXJpZCI6MSwidXNlcm5hbWUiOiJqb2huIn0.F8JdDOfmHwL-Z7TnaYBnp5fpaEW-Vn1obP4SALwA52o"
token2="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTc4NDMxOTIsInVzZXJpZCI6MiwidXNlcm5hbWUiOiJ0b20ifQ.ESH5xGXIyjPIF3lQHvpzl5faOIgwIYx9E4PbSn6Gh38"




#创建文章, jwt token非法 
curl -X POST http://localhost:8080/api/post\
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${token1}1111" \
  -d '{
    "title": "故乡的秋天",
    "content": "xxxxxx"
  }'
echo ""

#创建文章
curl -X POST http://localhost:8080/api/post\
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${token1}" \
  -d '{
    "title": "故乡的秋天",
    "content": "xxxxxx"
  }'
echo ""

curl -X POST http://localhost:8080/api/post\
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${token2}" \
  -d '{
    "title": "奇幻漂流记",
    "content": "xxxxxx"
  }'
echo ""

#查找所有文章
curl  http://localhost:8080/api/post\
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${token1}" 
echo ""

#查找某篇文章
curl  http://localhost:8080/api/post/1\
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${token1}" 
echo ""

#删除某篇文章
curl -X DELETE  http://localhost:8080/api/post/1\
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${token1}" 
echo ""

#更新某篇文章
curl -X PUT http://localhost:8080/api/post/2\
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${token2}" \
  -d '{
    "title": "奇幻漂流记",
    "content": "更新这个内容，现在是新内容"
  }'
echo ""

#给文章2添加评论
curl -X POST http://localhost:8080/api/comment\
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${token2}" \
  -d '{
    "postid" : 2,
    "content": "文章写得很好，希望博主多多更新"
  }'
echo ""

#查询文章2的评论
curl http://localhost:8080/api/post/2/comment\
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${token2}"
echo ""
