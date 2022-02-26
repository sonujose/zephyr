FROM golang:1.17-alpine3.15 AS builder

WORKDIR /go/src/app

COPY . .

# Generate/Updates swagger docs as per latest API changes
#RUN go get -u github.com/swaggo/swag/cmd/swag
#RUN swag init

# RUN go get

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o kube-spectrum .

FROM alpine:3.14.3 
RUN apk --no-cache add ca-certificates

# Copy local kubeconfig to local container
RUN mkdir /root/.kube
COPY kubeconfig /root/.kube/config

WORKDIR /usr/app/

COPY --from=builder /go/src/app .

EXPOSE 8080

ENTRYPOINT ["./kube-spectrum"]