# Stage 1 - wkhtmltopdf dependencies
FROM surnet/alpine-wkhtmltopdf:3.20.0-0.12.6-full AS wkhtmltopdf

# Stage 2 - Base image for Go, using Go 1.24.2 or compatible version
FROM golang:1.24-alpine3.20

# Install Go 1.24.2 manually (this may not be necessary if using official golang:1.24-alpine)
RUN apk add --no-cache \
    wget \
    && wget https://golang.org/dl/go1.24.2.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.24.2.linux-amd64.tar.gz \
    && rm go1.24.2.linux-amd64.tar.gz \
    && ln -s /usr/local/go/bin/go /usr/bin/go

# Add wkhtmltopdf required packages
RUN apk add --no-cache \
    libstdc++ libx11 libxrender \
    libxext fontconfig freetype \
    ttf-droid ttf-freefont ttf-liberation \
    bash ca-certificates wget

# Copy the wkhtmltopdf binary from "wkhtmltopdf" reference image
COPY --from=wkhtmltopdf /bin/wkhtmltopdf    /usr/local/bin/wkhtmltopdf
COPY --from=wkhtmltopdf /bin/wkhtmltoimage  /usr/local/bin/wkhtmltoimage
COPY --from=wkhtmltopdf /bin/libwkhtmltox*  /usr/local/bin

# Ensure binaries are executable
RUN chmod +x /usr/local/bin/wkhtmltopdf && \
    chmod +x /usr/local/bin/wkhtmltoimage

# Add /usr/local/bin to PATH explicitly
ENV PATH="/usr/local/bin:${PATH}"

# Install Go Air autoreload package
RUN go install github.com/air-verse/air@v1.52.3

# Install MySQL client for database interaction
RUN apk add --no-cache mysql-client

# Set working directory to /app
WORKDIR /app

# Copy application code into the container
COPY . . 

# Ensure go modules are tidy
RUN go mod tidy

# Expose port 3000
EXPOSE 3000

# Run autoreload or direct binary depending on environment
CMD ["air", "-c", ".air.toml"]
