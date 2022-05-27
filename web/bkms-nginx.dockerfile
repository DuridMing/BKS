FROM nginx:alpine
WORKDIR /app
COPY dist /usr/share/nginx/html
COPY nginxconf/ssl.csr /etc/nginx/cert/ssl.csr
COPY nginxconf/ssl.key /etc/nginx/cert/ssl.key
RUN rm /etc/nginx/conf.d/default.conf
COPY nginxconf/nginx.conf /etc/nginx/conf.d/

EXPOSE 443