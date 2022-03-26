FROM golang:1.17 as go_builder
ENV GOSUMDB="off" \
    CGO_ENABLED=0
WORKDIR "/build"
COPY go.mod go.sum ./
RUN go mod download -x
ADD ./go_application/* ./
RUN go build -mod readonly -o server \
    && chmod +x server

FROM busybox:stable as production
USER nobody
WORKDIR "/opt/app"
EXPOSE 8080
ADD config/mwh-config.json .
COPY --from=go_builder /build/server .
CMD ["./server", "-config=mwh-config.json"]
