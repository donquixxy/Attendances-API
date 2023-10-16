# builder image
FROM golang:1.21.3-alpine as builder
WORKDIR /build
COPY . .
RUN GOOS=linux 
RUN go build -o dgtest .

# generate clean, final image for end users
FROM alpine
RUN apk add --no-cache curl
RUN apk update && apk add ca-certificates && apk add tzdata && apk add git
COPY --from=builder /build .
ENV TZ="Asia/Makassar"
EXPOSE 9087

CMD ./dgtest