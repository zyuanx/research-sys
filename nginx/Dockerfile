FROM nginx:latest

COPY ./nginx.conf /etc/nginx/conf.d/

RUN rm /etc/nginx/conf.d/default.conf \
 && mkdir -p /usr/share/nginx/html/static \
 && mkdir -p /usr/share/nginx/html/media \
 && mkdir -p /usr/share/nginx/ssl

CMD ["nginx", "-g", "daemon off;"]
