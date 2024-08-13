# syntax=docker/dockerfile:1

# Frontend
FROM node:22-alpine3.19 AS frontend

WORKDIR /app

COPY frontend/yarn.lock frontend/package.json ./
RUN yarn install

COPY frontend ./
RUN yarn build

# Backend
FROM golang:1.22-alpine3.19

COPY --from=frontend /app/dist /app/frontend/dist
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /vue-go-playground main/server.go

EXPOSE 3000

# Run
CMD ["/vue-go-playground"]

# Run from root folder:
# docker build -t insert_name -f build/package/Dockerfile .
# docker run -p 3000:3000 insert_name