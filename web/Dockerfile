FROM debian:10-slim

RUN apt update && apt -y install \
    curl \
    nmap \
    netcat

WORKDIR /security
COPY entrypoint.sh ./entrypoint.sh
RUN chmod +x entrypoint.sh 

ENTRYPOINT ["./entrypoint.sh"]