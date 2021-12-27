FROM golang:1.16-alpine as build
WORKDIR /build
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download -x
COPY src ./src
RUN CGO_ENABLED=0 go build -o /build/server -v ./src

FROM scratch
WORKDIR /run
COPY --from=build /build/server .
ENTRYPOINT [ "/run/server" ]
