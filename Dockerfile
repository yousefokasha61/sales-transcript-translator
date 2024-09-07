FROM golang:1.23-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Youssef Okasha <yousefokasha61@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

ENV HTTP_PORT 8080
ENV LOG_LEVEL InfoLevel

# Expose port 8080 to the outside world
EXPOSE 8080

# Build the Go app
RUN go build -o main .


# Run the executable
CMD ["./main"]