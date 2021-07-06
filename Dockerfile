# waybill_center
#
# VERSION               1.0.1
FROM golang:1.15.5-alpine3.12 as builder1
COPY . /waybill_center
WORKDIR /waybill_center
#RUN ls && cd cmd/config && ls

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GOPRIVATE=github.com/hzlpypy
ENV GOMOD=/waybill_center/go.mod

COPY cmd/main.go .
COPY /docker/.ssh /root/.ssh
COPY /docker/.gitconfig /root/.gitconfig
RUN sed -i 's:dl-cdn.alpinelinux.org:mirrors.tuna.tsinghua.edu.cn:g' /etc/apk/repositories && apk add git openssh-client
RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"
RUN go mod download &&  go build -o ./main .
RUN mkdir /usr/local/waybill_center && mv main /usr/local/waybill_center
WORKDIR /usr/local/waybill_center
RUN mkdir log
CMD ["./main"]