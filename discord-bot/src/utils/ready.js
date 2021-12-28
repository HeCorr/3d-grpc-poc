function nodeLogger() {
    const nodelogger = require('phx-nodelogger')

    const logger = new nodelogger()
    logger.phxLogger(
    `gRPC`, //header
    '500', // headerWidth
    'red', // headerColor
    `Discord.js files by Phoenix#8033\n\ngRPC Generation by Nino#5624`, //body
    'disabled', // backgroundColor
    'blue', // borderColor
    'single', //borderStyle
    false)} //fullBorders

    module.exports = { nodeLogger }