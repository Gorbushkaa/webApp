FROM node:12-alpine as builder

RUN mkdir /app
WORKDIR /app

COPY . .

RUN yarn install

RUN yarn run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

