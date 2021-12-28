function nodeLogger() {
    const nodelogger = require('phx-nodelogger')

    const logger = new nodelogger()
    logger.phxLogger(
    `gRPC`, //header
    '500', // headerWidth
    'red', // headerColor
    `Discord.js files by Phoenix557\n\ngRPC Generation by HeCorr`, //body
    'disabled', // backgroundColor
    'blue', // borderColor
    'single', //borderStyle
    false)} //fullBorders

    module.exports = { nodeLogger }