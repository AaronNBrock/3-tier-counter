FROM node:14.9-alpine as builder
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install
COPY . .
RUN yarn build

FROM nginx:1.19-alpine
WORKDIR /app
COPY --from=builder /app/dist /app
COPY create-config-js.sh /docker-entrypoint.d/
COPY nginx.conf /etc/nginx/nginx.conf