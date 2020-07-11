FROM golang

WORKDIR /go/src/app

COPY . .
RUN go build
CMD ["./obliwonk"]