const { nodeLogger } = require('./utils/ready.js')
const { requestRender } = require('./grpc.js')
const { Client, Intents } = require('discord.js')
const { validateInstruction, parseInstruction } = require('./parser.js')

// Discord client
const client = new Client({ intents: [Intents.FLAGS.DIRECT_MESSAGES, Intents.FLAGS.GUILDS, Intents.FLAGS.GUILD_MESSAGES], partials: ['CHANNEL'] })

// on bot ready, prints banner in console
client.on('ready', () => {
    nodeLogger()
})

// listen to messages
client.on('messageCreate', async (message) => {
    // ignore pings from bots
    if (message.author.bot) return false

    // ignore @everyone and @here
    if (message.content.includes("@here") || message.content.includes("@everyone")) return false

    // ignoring replied messages and only responding to direct pings ("@gRPC")
    if (message.mentions.has(client.user.id) || message.channel.type == 'DM') {
        const msg = message.content.replace(/<@.*?>/gm, '').trim()

        // check for empty or 'help' message
        if (['', 'help'].includes(msg)) {
            if (message.channel.type == 'DM') {
                await message.reply({ content: "Send me a list of shapes to draw with their position and scale, one on each line.\nSupported shapes: `cube`, `sphere`, `plane`, `cylinder`, `icosahedron`.\nSyntax: `shape(xPos, yPos, zPos)` or `shape(xPos, yPos, zPos, xScale, yScale, zScale)`.\nExample (cube with a cone on top):\n```cube(0, 0, 0)\ncone(0, 0, 1, 0.5, 0.5, 1)\n```" })
            } else {
                await message.reply({ content: "Ping me including a list of shapes to draw with their position and scale, one on each line.\nSupported shapes: `cube`, `sphere`, `plane`, `cylinder`, `icosahedron`.\nSyntax: `shape(xPos, yPos, zPos)` or `shape(xPos, yPos, zPos, xScale, yScale, zScale)`.\nExample (cube with a cone on top):\n```cube(0, 0, 0)\ncone(0, 0, 1, 0.5, 0.5, 1)\n```" })
            }
            return
        }
        try {
            // extract instructions from message (one per line)
            const instructions = msg.split('\n')
            let instrQueue = []

            instructions.forEach((inst, i) => {
                // ignore instructions starting with // or #
                if (inst.startsWith("//") || inst.startsWith("#")) {
                    return
                }
                // if the instruction is invalid, this function throws.
                validateInstruction(inst, i + 1)
                // parse and push instruction to queue
                instrQueue.push(parseInstruction(inst))
            })

            // yes, I know this looks weird. no, I don't care. it works.
            const [imageBytes, renderTime] = await requestRender(instrQueue)
            const imageData = Buffer.from(imageBytes)

            message.reply({
                content: '`Render time: ' + renderTime + '`',
                files: [
                    { attachment: imageData, name: "render.png" }
                ]
            })
        } catch (e) {
            console.error(e)
            if (message.channel.type == 'DM') {
                await message.reply({ content: `Error: ${e}.\nTry saying 'help'.` })
            } else {
                await message.reply({ content: `Error: ${e}.\nPing me without saying anything else to obtain more info.` })
            }
        }
    }
})

// bot login
client.login(process.env.DISCORD_BOT_TOKEN).catch(e => {
    console.log(`Failed to login: ${e}`)
})
