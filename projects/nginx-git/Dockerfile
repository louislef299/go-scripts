FROM teamfruit/nginx-fcgiwrap:latest

# Run updates
RUN apt update && apt install git fcgiwrap apache2-utils nginx -y

# Create users
# RUN useradd -m -s /bin/bash -g root -G sudo -u 1000 user 
# RUN  echo 'user:foo' | chpasswd

RUN mkdir -p /srv/git

# Create the repository
WORKDIR /srv/git
RUN mkdir repo1
WORKDIR /srv/git/repo1

# Setup git server
RUN git init . --bare --shared
RUN git update-server-info
RUN git config --local http.receivepack true
RUN git config --bool core.bare true

# Give that directory the proper ownership for nginx
RUN chown -R www-data:www-data /srv/git
RUN chmod -R 755 /srv/git

COPY default.conf /etc/nginx/conf.d/.
RUN htpasswd -bc /etc/nginx/.gitpasswd user foo

# docker run -dp 32000:80 -v brew-lbctl-git:/srv/git/repo1 nginx-git
# git clone http://user:foo@127.0.0.1:32000/repo1