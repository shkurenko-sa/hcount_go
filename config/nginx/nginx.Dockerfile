FROM nginx:1.21-alpine

RUN rm /etc/nginx/conf.d/default.conf
COPY ./config/nginx/nginx.conf /etc/nginx/conf.d/