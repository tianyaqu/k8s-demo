FROM iron/go:dev as builder

RUN go get github.com/julienschmidt/httprouter && go get github.com/BurntSushi/toml && go get github.com/gomodule/redigo/redis
WORKDIR /go/src/app
COPY *.go src/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api-server src/*.go

FROM iron/go
WORKDIR /app
# Now just add the binary
COPY --from=builder /go/src/app/api-server .
ENTRYPOINT ["./api-server"]
