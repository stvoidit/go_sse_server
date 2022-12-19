# build golang backend application
FROM golang:1.19-alpine AS backend
WORKDIR /app
COPY src/backend .
ENV CGO_ENABLED=0 GOARCH=amd64 GOOS=linux GOFLAGS=-mod=vendor
RUN go build -ldflags "-s -w" -o ./build/go_sse_server ./main.go

# build vue frontend static files
FROM node:18-alpine AS frontend
WORKDIR /app
RUN npm install -g pnpm
COPY src/frontend/package.json src/frontend/pnpm-lock.yaml src/frontend/vite.config.ts src/frontend/tsconfig.json src/frontend/index.html ./
RUN pnpm install
COPY src/frontend/public public
COPY src/frontend/src src
RUN NODE_ENV=production pnpm build

# final stage, full build application
FROM alpine:latest AS backend-build
WORKDIR /build
COPY --from=frontend app/dist ./static
COPY --from=backend app/build .
ENTRYPOINT [ "./go_sse_server" ]
