FROM golang:alpine3.16

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code
COPY cmd ./cmd
COPY internal ./internal

# Build
RUN go build -o pokeapp ./cmd
RUN chmod +x pokeapp

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
ENV PORT=8080
# Run
CMD [ "/app/pokeapp" ]