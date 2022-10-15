FROM nginx:latest

# Run updates
RUN apt update && apt install git fcgiwrap apache2-utils nginx -y

RUN mkdir -p /srv/git

# Give that directory the proper ownership for nginx
RUN chown -R www-data. /srv/git

# Create the repository
WORKDIR /srv/git
RUN mkdir repo1
WORKDIR /srv/git/repo1
RUN git init . --bare --shared
RUN git update-server-info
COPY config /srv/git/repo1/.

RUN git config --bool core.bare true

RUN chown -R www-data:www-data /srv/git

COPY default.conf /etc/nginx/conf.d/.
RUN htpasswd -bc /etc/nginx/.gitpasswd user foo