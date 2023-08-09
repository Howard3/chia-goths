# Step 1: Build the UI
FROM node:18 as build-ui
WORKDIR /app

COPY . .
RUN npm install
RUN npx tailwindcss -i ./src/css/main.css -o ./assets/css/main.css


# Step 2: Build the executable
FROM golang:1.20 AS build
WORKDIR /app

COPY . .
COPY --from=build-ui /app/assets/css/main.css ./assets/css/main.css

ENV CGO_ENABLED=0

RUN go build -o chia-goths .

# Step 3: Create a minimal final image
FROM debian:bullseye-slim

COPY --from=build /app/chia-goths /usr/local/bin/chia-goths
ENV LISTEN_ADDR=0.0.0.0:8080

ENTRYPOINT ["chia-goths"]
