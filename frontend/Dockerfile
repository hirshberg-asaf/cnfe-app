FROM golang AS builder

WORKDIR /build

# Let's cache modules retrieval - those don't change so often
# Copy the code necessary to build the application
# You may want to change this to copy only what you actually need.
COPY main.go .

# Build the application
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/prometheus/client_golang/prometheus/promauto
RUN go get github.com/prometheus/client_golang/prometheus/promhttp
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -o /program ./main.go


# Let's create a /dist folder containing just the files necessary for runtime.
# Create the minimal runtime image
FROM alpine
EXPOSE 80 9000
COPY --from=builder /program /program
COPY static/ ./static/

CMD ["/program"]
