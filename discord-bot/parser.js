const validateRegex = /^[a-zA-z]+\([+-]?([0-9]*[.])?[0-9]+, ?[+-]?([0-9]*[.])?[0-9]+, ?[+-]?([0-9]*[.])?[0-9]+, ?[+-]?([0-9]*[.])?[0-9]+, ?[+-]?([0-9]*[.])?[0-9]+, ?[+-]?([0-9]*[.])?[0-9]+\)$/
const parseRegex = /^(\w+)\((.*)\)/

function validateInstruction(inst, line) {
    if (!validateRegex.test(inst)) {
        throw `Invalid instruction on line ${line}`
    }
}

function parseInstruction(inst) {
    const matches = parseRegex.exec(inst)
    const floatValues = matches[2].replace(/ +/, '').split(',')
    return {
        shape: matches[1],
        position: {
            x: floatValues[0],
            y: floatValues[1],
            z: floatValues[2],
        },
        scale: {
            x: floatValues[3],
            y: floatValues[4],
            z: floatValues[5],
        }
    }
}

module.exports = { validateInstruction, parseInstruction };
