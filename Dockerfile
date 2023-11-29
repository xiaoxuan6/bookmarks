FROM golang:1.20.4-alpine3.18 AS build

COPY . /src
WORKDIR /src

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod tidy

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o bookmarks .

FROM alpine

COPY --from=build /src/bookmarks /app/bookmarks
COPY --from=build /src/.env /app/.env
COPY --from=build /src/data /app/data

WORKDIR /app

EXPOSE 8080

RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

ENTRYPOINT ["./bookmarks"]


