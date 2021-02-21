FROM  golang:1.15.8 AS build-stage
WORKDIR /app/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build 

FROM frolvlad/alpine-glibc
WORKDIR /app/
COPY --from=build-stage /app/forum-back .
COPY ./ssl/forumWebApi.crt ./ssl/forumWebApi.key  ./ssl/

EXPOSE 8080:8080
CMD ["./forum-back"]