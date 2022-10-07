FROM ubuntu:latest

# Run updates
RUN apt update && apt install  openssh-server sudo -y

# Add user to container
# -d is the home directory for the new user
# -u is a specific user id
# -g is a specific group id
# -s is the login shell
RUN useradd -rm -d /home/ubuntu -s /bin/bash -g root -G sudo -u 1000 test 

# change password to test
RUN  echo 'test:test' | chpasswd

RUN service ssh start
EXPOSE 22

CMD ["/usr/sbin/sshd","-D"]
# docker run -dp 80:22 ssh-server