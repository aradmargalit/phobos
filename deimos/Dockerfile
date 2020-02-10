FROM node:latest as builder

ARG API_URL
ENV API_URL=$API_URL

WORKDIR /usr/src/app
COPY package.json yarn.lock ./
RUN yarn
COPY . ./
RUN yarn build

FROM nginx:1.12-alpine
COPY --from=builder /usr/src/app/build /usr/share/nginx/html
EXPOSE 80

# We want to use our own nginx configuration file
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d

CMD ["nginx", "-g", "daemon off;"]