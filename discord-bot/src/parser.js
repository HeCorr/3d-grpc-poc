const validateRegex3 = /^[a-zA-z]+\([+-]?([0-9]*[.])?[0-9]+, ?[+-]?([0-9]*[.])?[0-9]+, ?[+-]?([0-9]*[.])?[0-9]+\);?/
const validateRegex6 = /^[a-zA-z]+\([+-]?([0-9]*[.])?[0-9]+, ?[+-]?([0-9]*[.])?[0-9]+, ?[+-]?([0-9]*[.])?[0-9]+, ?[+-]?([0-9]*[.])?[0-9]+, ?[+-]?([0-9]*[.])?[0-9]+, ?[+-]?([0-9]*[.])?[0-9]+\);?/
const parseRegex = /^(\w+)\((.*)\)/

function validateInstruction(inst, line) {
    if (!validateRegex3.test(inst) && !validateRegex6.test(inst)) {
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
            x: floatValues.length == 6 ? floatValues[3] : 1,
            y: floatValues.length == 6 ? floatValues[4] : 1,
            z: floatValues.length == 6 ? floatValues[5] : 1,
        }
    }
}

module.exports = { validateInstruction, parseInstruction }
