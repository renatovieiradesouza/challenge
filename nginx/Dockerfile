FROM nginx:1.21-alpine

COPY nginx.conf /etc/nginx/nginx.conf

RUN mkdir -p /var/cache/nginx/temp

EXPOSE 8081

CMD ["nginx", "-g", "daemon off;"]

