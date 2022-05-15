FROM golang:1.17.9

ENV NODE=production
ENV PORT=8080
ENV CONTEXT_PATH=/course-listenner
ENV PATH_STRAPI_ARTICLE=/api/courses
ENV HOST_STRAPI_SERVICE=http://strapi-information-gateway:8080
ENV PATH_STRAPI_INFORMATION_GATEWAY=/strapi-information-gateway/api/strapi

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

EXPOSE 8080

CMD ["app"]