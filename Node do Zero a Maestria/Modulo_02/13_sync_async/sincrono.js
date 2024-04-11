const fs = require('fs');

console.log("Início");
fs.writeFileSync("sync.txt", "Texto para ser colado");
console.log("Termino");