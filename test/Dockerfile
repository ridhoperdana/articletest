FROM node:10-alpine3.11

RUN apk update && apk upgrade && \
    apk --no-cache --update add gcc musl-dev ca-certificates curl && \
    rm -rf /var/cache/apk/*

RUN npm install -g newman

WORKDIR /app

CMD ["newman", "run", "postman.json"]
