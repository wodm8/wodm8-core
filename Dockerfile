FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/wodm8/wodm8-core
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/wodm8-core-api main.go

#Building image with binary
FROM scratch
COPY --from=build /go/bin/wodm8-core-api /go/bin/wodm8-core-api
ENTRYPOINT ["/go/bin/wodm8-core-api"]