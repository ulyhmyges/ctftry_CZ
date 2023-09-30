FROM golang:1.21.1-alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR /go/ctf_api
COPY . .

RUN go install
RUN go build -o /go/myapp/bin/ctf_api

FROM scratch
COPY --from=builder /go/myapp/bin/ctf_api /go/bin/ctf_api
ENTRYPOINT ["/go/bin/ctf_api"]