FROM node:20.19.4-bullseye-slim
MAINTAINER Christian Lück <christian@lueck.tv>

RUN npm install -g json-server

WORKDIR /data
COPY db.json /data/db.json

EXPOSE 80
ADD run.sh /run.sh
ENTRYPOINT ["bash", "/run.sh"]
CMD []