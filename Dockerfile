#builder
FROM golang:alpine as builder
WORKDIR /home
COPY . .
RUN go build -o restapps-service main.go

#final image
FROM alpine
RUN apk add tzdata
COPY --from=builder /home/restapps-service .
EXPOSE 5005
CMD ./restapps-service