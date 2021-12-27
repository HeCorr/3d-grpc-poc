const { nodeLogger } = require('./utils/ready.js');
const { requestRender } = require('./grpc.js');
const { Client, Intents } = require('discord.js');
const { validateInstruction, parseInstruction } = require('./parser.js');

const client = new Client({ intents: [Intents.FLAGS.DIRECT_MESSAGES, Intents.FLAGS.GUILDS, Intents.FLAGS.GUILD_MESSAGES], partials: ['CHANNEL'] });

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
    if (message.mentions.has(client.user.id) || message.channel.type == 'DM') {
        const msg = message.content.replace(/<@.*?>/gm, '').trim()

        if (['', 'help'].includes(msg)/* message.content.trim().startsWith('<') && message.content.trim().endsWith('>') */) {
            if (message.channel.type == 'DM') {
                await message.reply({ content: "Send me a list of shapes to draw with their position and scale, one on each line.\nSupported shapes: `cube`, `sphere`, `plane`, `cylinder`, `icosahedron`.\nSyntax: `shape(xPos, yPos, zPos)` or `shape(xPos, yPos, zPos, xScale, yScale, zScale)`.\nExample (cube with a cone on top):\n```cube(0, 0, 0)\ncone(0, 0, 1, 0.5, 0.5, 1)\n```" })
            } else {
                await message.reply({ content: "Ping me including a list of shapes to draw with their position and scale, one on each line.\nSupported shapes: `cube`, `sphere`, `plane`, `cylinder`, `icosahedron`.\nSyntax: `shape(xPos, yPos, zPos)` or `shape(xPos, yPos, zPos, xScale, yScale, zScale)`.\nExample (cube with a cone on top):\n```cube(0, 0, 0)\ncone(0, 0, 1, 0.5, 0.5, 1)\n```" })
            }
            return
        }
        try {
            const instructions = msg.split('\n')
            let instrQueue = []

            instructions.forEach((inst, i) => {
                validateInstruction(inst, i + 1);
                instrQueue.push(parseInstruction(inst))
            })

            const imageBytes = await requestRender(instrQueue)
            const imageData = Buffer.from(imageBytes);

            message.reply({
                files: [
                    { attachment: imageData, name: "render.png" }
                ]
            })
        } catch (e) {
            console.error(e);
            if (message.channel.type == 'DM') {
                await message.reply({ content: `Error: ${e}.\nTry saying 'help'.` })
            } else {
                await message.reply({ content: `Error: ${e}.\nPing me without saying anything else to obtain more info.` })
            }
        }
    }
})

// iinitialize bot
client.login(process.env.DISCORD_BOT_TOKEN).catch(e => {
    console.log(`Failed to login: ${e}`)
})
