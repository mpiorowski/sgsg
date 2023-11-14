apt-get install certbot python3-certbot-nginx

sudo certbot certonly --nginx -d www.sgsg.bearbyte.org -d sgsg.bearbyte.org

sudo certbot renew --dry-run

sudo systemctl disable nginx
