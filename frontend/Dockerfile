FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./

RUN npm install --no-audit --no-fund && npm cache clean --force

COPY . .

ARG VUE_APP_API_URL=/api
ENV VUE_APP_API_URL=$VUE_APP_API_URL
ENV NODE_OPTIONS=--openssl-legacy-provider

RUN npm run build

FROM nginx:1.27-alpine

RUN apk add --no-cache curl

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

USER 101:101
EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]