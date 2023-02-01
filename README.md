# ffmpeg-api-tool
## 1. 构建镜像

```shell
 docker build -t ffmpeg_api_tool ./
```

## 2. 创建容器

```shell
docker run -itd --name=ffmpeg_api_tool -p 8099:8080 --entrypoint='bash' ffmpeg_api_tool
```

## 3. Api文档

```
http://127.0.0.1:8099/swagger/index.html
```

