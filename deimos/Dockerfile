###########
# PHASE 1 #
###########
FROM node:latest as builder

WORKDIR /usr/src/app
COPY package.json yarn.lock ./
RUN yarn
COPY . ./

# Build args, usually passed through the docker-compose
ARG REACT_APP_API_URL

# Set environment variables
ENV REACT_APP_API_URL=$REACT_APP_API_URL

# Finally, build the app! Production edition!
RUN yarn build


###########
# PHASE 2 #
###########
FROM nginx:1.12-alpine
COPY --from=builder /usr/src/app/build /usr/share/nginx/html

# We want to use our own nginx configuration file
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d

# Default port exposure
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]