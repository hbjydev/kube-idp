FROM golang:1.18-alpine AS build

RUN apk add make gcc libc-dev

WORKDIR /build

COPY . .

RUN go mod tidy
RUN make

FROM scratch

WORKDIR /

COPY --from=build /build/dist/server .

ENTRYPOINT [ "/server" ]
