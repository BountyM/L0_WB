ARG GO_IMAGE
ARG GO_IMAGE_VERSION

FROM ${GO_IMAGE}:${GO_IMAGE_VERSION}


WORKDIR /usr/src/app

COPY . .
RUN go mod download
RUN go build -o server ./cmd/main.go

CMD [ "./server" ]

