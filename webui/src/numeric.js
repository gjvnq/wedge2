export default {
  parse (val) {
    var re = /^([-0-9']*)(?:|(?:.|,)([0-9]*))$/
    var match = val.toString().match(re)
    if (match === null) {
      return 0 / 0
    }
    var intPart = match[1].replace(/'/g, '')
    var decPart = match[2]
    // Avoid problems
    if (intPart === undefined) {
      intPart = ''
    }
    if (decPart === undefined) {
      decPart = ''
    }
    // Pad with zeros if necessary
    while (decPart.length < 8) {
      decPart += 0
    }
    // Exclude zeros if necessary
    decPart = decPart.substr(0, 8)
    // Convert to the integer form
    var final = Number(intPart + decPart)
    return final
  },
  format (int) {
    var str = int.toString()
    // Take care of non-numbers
    if (!isFinite(int)) {
      return str
    }
    // Zero pad just in case
    if (str.length < 9) {
      str = '0'.repeat(9 - str.length) + str
    }
    // Get parts
    var intPart = str.substr(0, str.length - 8)
    var decPart = str.substr(str.length - 8, str.length)
    // Remove uncessary zeros making sure to keep at least two decimal places if decPart is NOT zero
    if (Number(decPart) === 0) {
      return intPart
    }
    decPart = decPart.split('')
    while (decPart.length > 2 && decPart[decPart.length - 1] === '0') {
      decPart.pop()
    }
    decPart = decPart.join('')
    return intPart + '.' + decPart
  }
}
