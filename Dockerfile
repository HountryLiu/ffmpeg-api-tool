FROM golang:1.18-buster AS builder
ENV GO111MODULE on
ENV CGO_ENABLED 1
ENV GOFLAGS -mod=vendor
COPY . /go/src
WORKDIR /go/src
RUN go build -v -o ffmpeg_api_tool -ldflags "-s -w" .

# ===========================

FROM jrottenberg/ffmpeg
COPY --from=builder /go/src/ffmpeg_api_tool /
WORKDIR /
ENV TZ Asia/Shanghai
ENV LOG_LEVEL 0
EXPOSE 8080
ENTRYPOINT ["sh","-c","/ffmpeg_api_tool"]

