# Instant Messager

![Tests](https://github.com/Echomo-Xinyu/Instant-Messager/actions/workflows/test.yml/badge.svg)

This is a Instant Messager system completed as the [backend assignment of 2023 TikTok Tech Immersion](https://bytedance.sg.feishu.cn/docx/P9kQdDkh5oqG37xVm5slN1Mrgle). In this assignment, I have implemented the `send`, `pull`in `./rpc-server/handler.go` and connect application to a redis database.

In order to run the application, you need to have `go`, `redis` and `docker` installed.

After running `docker compose -f "docker-compose.yml" up -d --build`, you will have the application running on docker.

`curl -X POST 'localhost:8080/api/send?sender=a&receiver=b&text=hi'` means that `a` sends `b` a message of content `hi`.

`curl 'localhost:8080/api/pull?chat=a%3Ab'` pulls the messages in chat between `a` and `b`.
