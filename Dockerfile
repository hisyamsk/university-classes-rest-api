FROM golang:1.19-bullseye AS build

RUN useradd -u 1001 nonroot

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main_app

##
FROM scratch

WORKDIR /app

COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /app/main_app main_app

USER nonroot

EXPOSE 8000

CMD [ "./main_app" ]