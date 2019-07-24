FROM golang:1.10 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR /go/src/github.com/moustafab/gitrepos
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /gitrepos .

FROM scratch

COPY --from=builder /gitrepos ./

ENTRYPOINT ["./gitrepos"]
