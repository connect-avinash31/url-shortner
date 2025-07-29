#base image for docker 
FROM golang:1.20-alpine

ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

#deifne the working directory
WORKDIR /app

# Build your Go app
RUN go build -o urlshortner

# now copy everythoing from the current directory to the working directory
COPY . .

CMD ["./urlshortner"]

