FROM node:20-alpine
WORKDIR /app
COPY does-not-exist.txt .
COPY server.js .
CMD ["node","server.js"]
