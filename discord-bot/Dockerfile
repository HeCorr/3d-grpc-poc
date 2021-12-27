FROM node:16-alpine
WORKDIR /bot
ENV ENV=docker
COPY package.json .
COPY yarn.lock .
RUN yarn install --production
COPY render_server.proto ./render_server.proto
COPY src ./src
ENTRYPOINT [ "node", "src/main.js" ]
