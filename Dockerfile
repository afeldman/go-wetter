FROM golang as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o weather .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /bin
COPY --from=builder /app/weather .

EXPOSE 2510

CMD ["./weather","-release"]
