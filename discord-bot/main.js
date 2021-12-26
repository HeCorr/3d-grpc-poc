const config = require('./config.js');
const { nodeLogger } = require('./utils/ready.js')
const { requestRender } = require('./grpc.js');
const { Client, Intents } = require('discord.js');

const client = new Client({ intents: [Intents.FLAGS.GUILDS, Intents.FLAGS.GUILD_MESSAGES] });

// on bot ready, logs banner
client.on('ready', () => {
    nodeLogger()
});

// listen to messages
client.on('messageCreate', async (message) => {
    // ignore pings from bots
    if (message.author.bot) return false;

    // ignore @everyone and @here
    if (message.content.includes("@here") || message.content.includes("@everyone")) return false

    // ignoring replied messages and only responding to direct pings ("@gRPC")
    if (message.mentions.has(client.user.id)) {
        if (message.content = "draw") {
            try {
                const imageBytes = await requestRender()
                const imageData = Buffer.from(imageBytes);
                message.reply({
                    files: [
                        { attachment: imageData, name: "render.png" }
                    ]
                })
            } catch (e) {
                console.error(e);
                message.reply({ content: "error: " + e })
            }

        }
    }
})

// iinitialize bot
client.login(config.token).catch(e => {
    console.log(`Failed to login: ${e}`)
})
