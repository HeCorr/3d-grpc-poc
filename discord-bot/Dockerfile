FROM node:16-alpine
WORKDIR /bot
COPY package.json .
COPY yarn.lock .
COPY src ./src
ENV ENV=docker
RUN yarn install --production
ENTRYPOINT [ "node", "src/main.js" ]
