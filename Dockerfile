FROM golang:latest as go_builder
WORKDIR /go/src/mvh
ARG build_dir="/opt/mvh_build"

RUN go get github.com/lib/pq \
           github.com/gorilla/mux

ADD ./go_application/* ./
RUN go build -o /go/bin/mvh mvh


FROM debian:stable-slim
LABEL mainteiner="Vasiliy Badaev <againDDM@gmail.com>"
ARG app_dir="/opt/mvh"
WORKDIR $app_dir
EXPOSE 8000
ADD mwh-config.json $app_dir
COPY --from=go_builder /go/bin/mvh $app_dir/mvh
CMD ["./mvh", "-config=mwh-config.json"]
