FROM golang:1.21-alpine AS build
WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev
#dependencies
COPY ["go.mod","go.sum","./"]
RUN go mod download

#build
COPY . ./
RUN go build -o ./bin/app ./main.go

FROM alpine AS runer
COPY --from=build /usr/local/src/bin/app /
CMD ["/app"]