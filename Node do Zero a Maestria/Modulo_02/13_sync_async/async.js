const fs = require("fs");

console.log("Antes");

fs.writeFile("async.txt", "texto async", (err) => {
  setTimeout(() => {
    console.log("Arquivo criado.");
  }, 1000);
});

console.log("Depois");
