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

function checkCardMatches(line) {
  let matches = 0;
  const [winningNums, cardNums] = parseCardNumbers(line);

  cardNums.forEach((cn) => {
    if (winningNums.includes(cn)) {
      matches++;
    }
  });

  return matches;
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
  pointSum += checkCardPoints(lines[i]);
}

console.log("Part 1 - Points Sum:", pointSum);

const cardCounts = Array(lines.length).fill(1);
let matches = 0;
let localMatches;

for (let i = 0; i < lines.length; i++) {
  matches = checkCardMatches(lines[i]);

  if (matches > 0) {
    for (let j = 0; j < cardCounts[i]; j++) {
      localMatches = matches;
      for (
        let copyIdx = i + 1;
        copyIdx < lines.length && localMatches > 0;
        copyIdx++
      ) {
        cardCounts[copyIdx]++;
        localMatches--;
      }
    }
  }
}

const totalCards = cardCounts.reduce((sum, c) => sum + c, 0);

console.log("Part 2 - Total Cards:", totalCards);
