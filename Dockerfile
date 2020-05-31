FROM  golang:1.14.3 AS build-stage

WORKDIR /app/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build 

FROM frolvlad/alpine-glibc
WORKDIR /app/
COPY --from=build-stage /app/forum .

EXPOSE 8080:8080
CMD ["./forum"]