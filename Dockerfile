FROM golang:latest

WORKDIR /web/vegan-map-api
COPY . .
#RUN go build
CMD ["./gin-vegan-map-api"]
