FROM debian:10-slim

# bash colors
RUN echo "PS1='\[\e]0;${debian_chroot:+($debian_chroot)}\u@\H: \w\a\]${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\H\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]\$ '" >> ~/.bashrc

RUN apt update && apt -y install \
    curl \
    nmap \
    tcpdump \
    socat

WORKDIR /security
COPY entrypoint.sh ./entrypoint.sh
RUN chmod +x entrypoint.sh 

ENTRYPOINT ["./entrypoint.sh"]