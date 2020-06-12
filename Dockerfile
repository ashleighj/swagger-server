ARG project=swagger-server

# ------ BUILD STAGE -------
FROM golang:alpine AS builder
ARG project

LABEL stage=builder

# Basic libraries
RUN apk add --no-cache alpine-sdk

WORKDIR /${project}

# Install package dependencies
RUN go get -d -v ./...

# Copy project files
COPY . .

# Build application
RUN go build -o ./${project}

# ------- FINAL STAGE -------
FROM alpine
ARG project
ENV entry=${project}

# Create non-root user for app
RUN adduser -D -g 'gouser' gouser && \
    mkdir -p /${project} && \
    chown -R gouser:gouser /${project}

USER gouser

# Copy files from builder

COPY --from=builder /${project}/docs/* ${project}/docs/
COPY --from=builder /${project}/${project} ${project}/

EXPOSE 8080
WORKDIR /${project}
ENTRYPOINT ./${entry}