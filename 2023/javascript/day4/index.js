const fs = require("fs");

function parseCardNumbers(cardStr) {
  const splitCard = cardStr.split(": ");
  let [winningNums, cardNums] = splitCard[1].split(" | ");
  winningNums = winningNums.split(" ").filter((n) => n);
  cardNums = cardNums.split(" ").filter((n) => n);
  return [winningNums, cardNums];
}

function checkCardPoints(line) {
  let points = 0;
  const [winningNums, cardNums] = parseCardNumbers(line);

  cardNums.forEach((cn) => {
    if (winningNums.includes(cn)) {
      if (points === 0) {
        points = 1;
      } else {
        points *= 2;
      }
    }
  });

  return points;
}

const filename =
  process.argv.includes("t") || process.argv.includes("-t")
    ? "input-test.txt"
    : "input.txt";

const lines = fs
  .readFileSync(__dirname + "/" + filename)
  .toString()
  .split("\n")
  .filter((l) => l);

let pointSum = 0;
for (let i = 0; i < lines.length; i++) {
  console.log(lines[i]);

  pointSum += checkCardPoints(lines[i]);
}

console.log("Part 1 - Points Sum:", pointSum);
