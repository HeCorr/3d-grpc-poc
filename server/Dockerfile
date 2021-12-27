FROM golang:1.16-alpine as build
ENV GO111MODULE=on
WORKDIR /build
COPY go.mod .
COPY go.sum .
COPY src ./src
RUN go mod download -x
RUN CGO_ENABLED=0 go build -o /build/server -v ./src

FROM scratch
WORKDIR /run
COPY --from=build /build/server .
ENTRYPOINT [ "/run/server" ]
