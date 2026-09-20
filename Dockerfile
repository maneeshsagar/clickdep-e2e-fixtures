# syntax=docker/dockerfile:1
ARG NODE_VERSION=20
FROM node:${NODE_VERSION}-alpine AS base
LABEL org.opencontainers.image.title="features"
ENV NODE_ENV=production E2E_FROM_DOCKERFILE=baked-in
WORKDIR /app
RUN apk add --no-cache tini \
 && addgroup -S app && adduser -S app -G app
COPY --chown=app:app server.js entrypoint.sh ./
RUN chmod +x entrypoint.sh
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=3s CMD wget -qO- http://127.0.0.1:8080/ || exit 1
USER app
ENTRYPOINT ["/sbin/tini","--","/app/entrypoint.sh"]
CMD ["node","server.js"]
