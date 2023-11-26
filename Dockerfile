FROM golang:alpine AS builder
ARG VERSION
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/ggermis/helm-util/
COPY . .
RUN go get -d -v
RUN go build -ldflags "-s -w -X github.com/ggermis/helm-util/pkg/helm_util/version.version=${VERSION}" -o "/go/bin/helm-util"

FROM golang:alpine
RUN apk update && apk add --no-cache bash helm
COPY config.yaml /config/charts.yaml
COPY --from=builder /go/bin/helm-util /go/bin/helm-util
ENTRYPOINT ["/go/bin/helm-util", "-c", "/config/charts.yaml"]