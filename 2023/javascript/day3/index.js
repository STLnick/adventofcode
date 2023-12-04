const fs = require("node:fs");

const gridChecks = [
  [-1, -1],
  [-1, 0],
  [-1, 1],
  [0, -1],
  [0, 0],
  [0, 1],
  [1, -1],
  [1, 0],
  [1, 1],
];
const numberRegex = /\d/;
const symbolRegex = /[^a-zA-Z0-9\s\.]+/g;

function extractFullInt(line, startX, endX) {
  while (startX >= 0 && numberRegex.test(line[startX])) {
    startX--;
  }
  startX++;

  while (endX < line.length && numberRegex.test(line[endX])) {
    endX++;
  }
  endX--;

  return [startX, endX];
}

function main() {
  console.log("-- Day 3 --");

  const filename =
    process.argv.includes("t") || process.argv.includes("-t")
      ? "input-test.txt"
      : "input.txt";

  const lines = fs
    .readFileSync(__dirname + "/" + filename)
    .toString()
    .split("\n");

  let char, endX, index, foundNums, line, num, prevNum, regexResults, startX;
  let gearRatioSum = 0;
  let sum = 0;

  for (let y = 0; y < lines.length; y++) {
    line = lines[y];

    if (line === "") {
      continue;
    }

    regexResults = symbolRegex.exec(line);

    while (regexResults !== null) {
      foundNums = [];
      char = regexResults[0];
      index = regexResults["index"];

      gridChecks.forEach(([dy, dx]) => {
        if (numberRegex.test(lines[y + dy][index + dx])) {
          [startX, endX] = extractFullInt(
            lines[y + dy],
            index + dx - 1,
            index + dx + 1,
          );

          num = parseInt(lines[y + dy].slice(startX, endX + 1));
          if (num !== prevNum) {
            sum += num;
            foundNums.push(num);
          }
          prevNum = num;
        }
      });

      if (foundNums.length === 2) {
        gearRatioSum += foundNums[0] * foundNums[1];
      }

      regexResults = symbolRegex.exec(line);
    }
  }

  /**LOG*/ console.log("Sum of Part Numbers:", sum);
  /**LOG*/ console.log("Sum of Gear Ratios:", gearRatioSum);
}

main();
