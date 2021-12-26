const config = require('./config.js');
const Discord = require('discord.js');
const { nodeLogger } = require('./utils/ready.js')

const { Client, Intents } = require('discord.js');

const client = new Client({ intents: [Intents.FLAGS.GUILDS, Intents.FLAGS.GUILD_MESSAGES] });

// on bot ready, logs banner
client.on('ready', () => {
    nodeLogger()
});

// listen to messages
client.on('messageCreate', (message) => {
    // ignore pings from bots
    if (message.author.bot) return false;

    // ignore @everyone and @here
    if (message.content.includes("@here") || message.content.includes("@everyone")) return false

    // ignoring replied messages and only responding to direct pings ("@gRPC")
    if (message.mentions.has(client.user.id)) {
        if (message.content = "draw") {
            message.channel.send('nino needs to add this')
        }
    }
})

// iinitialize bot
client.login(config.token)
    .catch(e => {
        console.log(`Failed to login: ${e}`)
    })
