FROM node:latest as builder

WORKDIR /usr/src/app
COPY package.json yarn.lock ./
RUN yarn
COPY . ./
RUN yarn build

FROM nginx:1.12-alpine
COPY --from=builder /usr/src/app/build /usr/share/nginx/html

# We want to use our own nginx configuration file
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d

# Default port exposure
EXPOSE 80

# Add bash
RUN apk add --no-cache bash

# Copy .env file and shell script to container
WORKDIR /usr/share/nginx/html
COPY ./env.sh .
COPY .env .

# Make our shell script executable
RUN chmod +x env.sh

# Start Nginx server
CMD ["/bin/bash", "-c", "/usr/share/nginx/html/env.sh && nginx -g \"daemon off;\""]