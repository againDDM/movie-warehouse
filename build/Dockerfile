FROM golang:1.17 as go_builder
ENV GOSUMDB="off" \
    CGO_ENABLED=0
WORKDIR "/build"
COPY go.mod go.sum ./
RUN go mod download -x
COPY . .
RUN go build -mod readonly -o server ./cmd/*.go \
    && chmod +x server

FROM busybox:stable as production
USER nobody
WORKDIR "/opt/app"
EXPOSE 8000
ADD config/mwh-config.json .
COPY --from=go_builder /build/server .
CMD ["./server", "runserver"]
HEALTHCHECK --interval=5s --timeout=2s --retries=1 CMD ./server ping
