FROM nginx:latest

COPY ./nginx.conf /etc/nginx/conf.d/

RUN rm /etc/nginx/conf.d/default.conf \
 && mkdir -p /usr/share/nginx/html/
CMD ["nginx", "-g", "daemon off;"]
