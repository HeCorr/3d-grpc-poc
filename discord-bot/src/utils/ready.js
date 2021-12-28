// Prints a banner in console on startup
function nodeLogger() {
    const nodelogger = require('hyperz-nodelogger')
    const logger = new nodelogger()
    // Entries: header, headerWidth, headerColor, body, backgroundColor, borderColor, borderStyle, fullBorders
    logger.hypelogger(`gRPC`, '500', 'red', `Discord.js bot by Phoenix557 \n\n gRPC Server by HeCorr`, 'disabled', 'blue', 'single', false)
}

module.exports = { nodeLogger }
