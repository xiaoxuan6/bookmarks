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

# .env 配置

```shell
title="xiaoxuan6‘s 的书签"
author="xiaoxuan6"
description="xiaoxuan6 的书签列表"
```

如何 `.env` 配置为空，将使用默认配置

```shell
title="xiaoxuan6's Bookmarks"
author="xiaoxuan6"
description="xiaoxuan6's Bookmarks"
```

# bookmarks.json 如何获取

工具 [tools](https://github.con/xiaoxuan6/tools) 中的 `bookmarks` 可以将浏览器书签导出为 `bookmarks.json` 文件