FROM golang:latest as builder
ADD . /app
WORKDIR /app/server
RUN go mod download
RUN go get -u github.com/pressly/goose/cmd/goose
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

# Build the React application
FROM node:alpine as node_builder
WORKDIR /webapp
COPY ./webapp/package.json .
RUN npm install
COPY ./webapp/ ./
RUN npm run build

# Final stage build, this will be the container
# that we will deploy to production
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
COPY --from=node_builder /webapp/build ./web
RUN chmod +x ./main
EXPOSE 8080
CMD ./main
