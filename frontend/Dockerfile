FROM node:16-alpine AS builder

WORKDIR /app

COPY package*.json ./

RUN npm install --no-audit --no-fund && npm cache clean --force

COPY . .

ARG VUE_APP_API_URL=http://89.169.156.255:8081
ENV VUE_APP_API_URL=$VUE_APP_API_URL

RUN npm run build

FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]