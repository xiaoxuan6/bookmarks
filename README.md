# bookmarks

将浏览器书签保存并运行到服务中

# Docker

```docker
docker run --name=bookmarks -p 8080:8080 -v .env:/app/.env -v /data:/app/data -d ghcr.io/xiaoxuan6/bookmarks:latest
```

* .env: 环境变量文件
* data: 书签文件夹

# Docker Compose

```yaml
version: '3.8'
services:
  bookmarks:
    image: ghcr.io/xiaoxuan6/bookmarks:latest
    container_name: bookmarks
    ports:
      - 8080:8080
    volumes:
      - $PWD/.env:/app/.env
      - $PWD/data:/app/data
    restart: always
```

# bookmarks.json 如何获取

工具 [tools](https://github.con/xiaoxuan6/tools) 中的 `bookmarks` 可以将浏览器书签导出为 `bookmarks.json` 文件