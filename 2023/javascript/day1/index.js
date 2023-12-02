const fs = require("node:fs");

const lookupTable = {
  one: 1,
  two: 2,
  three: 3,
  four: 4,
  five: 5,
  six: 6,
  seven: 7,
  eight: 8,
  nine: 9,
};

function getIntValue(val) {
  if (!isNaN(parseInt(val))) {
    return val;
  }

  if (Object.hasOwn(lookupTable, val)) {
    return lookupTable[val];
  }

  throw new Error(`getIntValue() :: couldn't find value for (${val})`);
}

function main() {
  let inputFileName = "input.txt";

  if (process.argv.includes("t") || process.argv.includes("-t")) {
    inputFileName = "input-test.txt";
  }

  const inputFile = fs.readFileSync(`${__dirname}/${inputFileName}`).toString();
  const lines = inputFile.split("\n");
  let idx, first, last, line, num, prevIdx;
  let sum = 0;
  let numEntries = [];

  for (let i = 0; i < lines.length; i++) {
    first = last = null;
    numEntries = [];
    line = lines[i];

    if (line === "") {
      continue;
    }

    // Check for integers in line
    for (let j = 0; j < line.length; j++) {
      num = parseInt(line[j]);

      if (!isNaN(num)) {
        numEntries.push({ value: num, index: j });
      }
    }

    // Check for number strings in line
    Object.keys(lookupTable).forEach((key) => {
      idx = prevIdx = line.indexOf(key);
      if (idx !== -1) {
        numEntries.push({ value: getIntValue(key), index: idx });
      }
      idx = line.lastIndexOf(key);
      if (idx !== -1 && idx !== prevIdx) {
        numEntries.push({ value: getIntValue(key), index: idx });
      }
    });

    // Determine first and last
    numEntries.forEach((entry) => {
      if (!first) {
        first = last = entry;
      }

      if (entry.index < first.index) {
        first = entry;
      }

      if (entry.index > last.index) {
        last = entry;
      }
    });

    sum += first.value * 10 + last.value;
  }

  console.log("Sum:", sum);
}

main();
