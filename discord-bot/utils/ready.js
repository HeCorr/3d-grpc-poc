function nodeLogger() {
    const nodelogger = require('hyperz-nodelogger')

    const logger = new nodelogger()
    // Entries: header, headerWidth, headerColor, body, backgroundColor, borderColor, borderStyle, fullBorders
    logger.hypelogger(`gRPC`, '500', 'red', `Discord.js files by Phoenix#8033 \n\n gRPC Generation by Nino#5624`, 'disabled', 'blue', 'single', false)}

    module.exports = {nodeLogger}