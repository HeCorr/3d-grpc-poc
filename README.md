# 3d-grpc-poc
A POC of a Discord.js bot that sends 3D rendering instructions to a Go server through [gRPC](https://grpc.io/) which responds with the image bytes which are then sent back on Discord.

## Project status
It's pretty much working, but I'd like to clean up the code some more and add more features.

It supports receiving instructions via DM as well as in guilds (by mentioning it).

It's currently being hosted 24/7 by me.

You can join my test server [here](https://discord.gg/bAbdu9p6rv) or invite the bot to your own guild with [this link](https://discord.com/api/oauth2/authorize?client_id=924427005053784084&permissions=34816&scope=bot).

## Libraries used
[Discord.js](https://discord.js.org/)

[Node](https://grpc.io/docs/languages/node/)'s and [Go](https://grpc.io/docs/languages/go/)'s official gRPC/Protobuf libraries

Fogleman's awesome [FauxGL renderer](https://github.com/fogleman/fauxgl)

## TODO's
- [x] Improve the code, make it more readable
- [x] Dockerize the server and the bot
- [x] Include render time in messages
