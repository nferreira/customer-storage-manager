FROM golang:1.14 as builder
COPY .netrc /root/.netrc
WORKDIR /go/customer-storage-manager
COPY . .
ENV GO111MODULE=on
WORKDIR cmd
RUN GOSUMDB=off CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o customer-storage-manager main.go
RUN cp customer-storage-manager /go/bin/customer-storage-manager

# Now copy it into our base image.
FROM alpine:latest
COPY --from=builder /go/bin/customer-storage-manager /
CMD ["/customer-storage-manager"]