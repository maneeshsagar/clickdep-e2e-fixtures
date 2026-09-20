FROM node:20-alpine
WORKDIR /app
COPY server.js .
CMD ["/does-not-exist"]
