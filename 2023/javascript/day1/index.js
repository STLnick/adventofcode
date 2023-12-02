const fs = require("node:fs");

function main() {
  let inputFileName = "input.txt";

  if (process.argv.includes("t") || process.argv.includes("-t")) {
    inputFileName = "input-test.txt";
  }

  const inputFile = fs.readFileSync(`${__dirname}/${inputFileName}`).toString();
  const lines = inputFile.split("\n");

  console.log({ inputFile, lines });

  let first, last, line, num;
  let sum = 0;

  for (let i = 0; i < lines.length; i++) {
    first = last = 0;
    line = lines[i];

    if (line === "") {
      continue;
    }

    for (let j = 0; j < line.length; j++) {
      num = parseInt(line[j]);

      if (!isNaN(num)) {
        if (!first) {
          first = last = num;
        } else {
          last = num;
        }
      }
    }

    console.log({ first, last });
    sum += first * 10 + last;
  }

  console.log("Sum:", sum);
}

main();
