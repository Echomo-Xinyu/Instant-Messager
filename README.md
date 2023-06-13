# Instant Messager

![Tests](https://github.com/Echomo-Xinyu/Instant-Messager/actions/workflows/test.yml/badge.svg)

This is a Instant Messager system completed as the [backend assignment of 2023 TikTok Tech Immersion](https://bytedance.sg.feishu.cn/docx/P9kQdDkh5oqG37xVm5slN1Mrgle). In this assignment, I have implemented the `send`, `pull`in `./rpc-server/handler.go` and connect application to a redis database.

In order to run the application, you need to have `go`, `redis` and `docker` installed.

After running `docker compose -f "docker-compose.yml" up -d --build`, you will have the application running on docker.

```
curl --location 'localhost:8080/api/send' \
--header 'Content-Type: application/json' \
--data '{
    "chat": "a1:a2",
    "text": "hi",
    "sender": "a2"
}'
```

means to send a message of content `hi` from `user: a2` in the chat between `user: a1` and `user: a2`;

```
curl --location 'localhost:8080/api/send' \
--header 'Content-Type: application/json' \
--data '{
    "chat": "a1:a2",
    "text": "test2",
    "sender": "a2"
}'
```

means to pull messages in the chat between `user: a1` and `user: a2`.
