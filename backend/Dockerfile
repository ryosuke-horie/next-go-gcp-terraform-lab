# デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.23.4-bullseye AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -buildvcs=false -trimpath -ldflags "-w -s" -o app 

# ---------------------------------------------------

FROM debian:bullseye-slim AS deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

EXPOSE 8080

CMD ["./app"]

# ---------------------------------------------------

FROM golang:1.23 AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]