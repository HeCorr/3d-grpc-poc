# 3d-grpc-poc
A POC of a Discord.js bot that sends 3D rendering requests to a Go server through gRPC and replies with the generated image


## Project status
It's pretty much working. It supports receiving instructions via DM as well as in guilds (by mentioning it).
Still not hosted permanently anywhere, but it's going to be soon.
If you wish to invite the bot, click [here](https://discord.com/api/oauth2/authorize?client_id=924427005053784084&permissions=34816&scope=bot).

## Libraries used
[Discord.js](https://discord.js.org/)
[Node](https://grpc.io/docs/languages/node/) and [Go](https://grpc.io/docs/languages/go/)'s official gRPC/Protobuf libraries
Fogleman's awesome [FauxGL renderer](https://github.com/fogleman/fauxgl)

## TODO's
- [ ] Add comments to the code
- [ ] Dockerize the server and the bot
- [x] Include render time in messages
