###########
# PHASE 1 #
###########



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