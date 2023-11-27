# bookmarks

将浏览器书签保存并运行到服务中

# Docker

```shell
docker run --name=bookmarks -p 8080:8080 -v .env:/app/.env -v bookmarks.json:/app/bookmarks.json -d ghcr.io/xiaoxuan6/bookmarks:latest
```

* .env: 环境变量文件
* bookmarks.json: 书签文件

# Docker Compose

```shell
version: '3.8'
services:
  bookmarks:
    image: ghcr.io/xiaoxuan6/bookmarks:latest
    container_name: bookmarks
    ports:
      - 8080:8080
    volumes:
      - $PWD/.env:/app/.env
      - $PWD/bookmarks.json:/app/bookmarks.json
    restart: always
```

# bookmarks.json 如何获取

工具 [tools](https://github.con/xiaoxuan6/tools) 中的 `bookmarks` 可以将浏览器书签导出为 `bookmarks.json` 文件