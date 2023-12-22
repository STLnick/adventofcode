const fs = require("fs");

function getLines() {
  let filename = "input.txt";
  if (process.argv.includes("t") || process.argv.includes("-t")) {
    filename = "input-test.txt";
  } else if (process.argv.includes("t2") || process.argv.includes("-t2")) {
    filename = "input-test2.txt";
  }

  return fs.readFileSync(__dirname + "/" + filename).toString().split("\n").filter(Boolean);
}

function part1(lines) {
  console.log("P1 :: lines");
  console.log(lines);

  return null;
}

function main() {
  const lines = getLines()
  const part1Result = part1(lines);
  console.log({ part1Result });
}

main();

