apt-get install certbot python3-certbot-nginx

sudo certbot certonly --nginx -d $1

sudo certbot renew --dry-run

sudo systemctl disable nginx
