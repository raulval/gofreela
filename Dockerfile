# Create Builder Image, to compile the source code into an executable
FROM golang:1.21.0 as builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# Create the final image, running the API and exposing port 8080
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/cmd/main .
ARG PORT
ENV PORT=$PORT
EXPOSE $PORT
CMD ["./main"]