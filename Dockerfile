FROM golang:1.16-buster AS build

# Copy our junk to the build directory.
RUN mkdir -p /build
WORKDIR /build
COPY . .

# Fail the build on purpose if there are no images.
RUN scripts/check-source-images.sh

# Install upx for binary compression.
RUN apt-get update
RUN apt-get install -y --no-install-recommends upx-ucl

# Build the binary and compress it.
RUN go build -o /bin/app -ldflags='-s' main.go
RUN upx -9 /bin/app

# The app image will be based on alpine.
FROM alpine AS app

RUN apk add --no-cache ca-certificates libc6-compat

WORKDIR /bin
COPY --from=build /bin/app .

ENV PORT=8080
EXPOSE ${PORT}

ENTRYPOINT [ "/bin/app" ]
