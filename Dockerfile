# https://codefresh.io/docs/docs/learn-by-example/golang/golang-hello-world/
FROM golang:1.18.1

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/git.sr.ht/~tsukii/pong

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Build the package
RUN go build ./...

# Install the package
RUN go install -v ./...

# This container exposes port 1323 to the outside world
EXPOSE 1323

# Run the executable
CMD ["pong"]
