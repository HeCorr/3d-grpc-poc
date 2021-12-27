FROM node:16-alpine
WORKDIR /bot
COPY package.json .
COPY yarn.lock .
COPY src ./src
ENV ENV=docker
COPY render_server.proto ./render_server.proto
RUN yarn install --production
ENTRYPOINT [ "node", "src/main.js" ]
