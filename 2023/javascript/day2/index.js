const fs = require("node:fs");

const cubes = {
  red: 12,
  green: 13,
  blue: 14,
};

const minCubes = {
  red: -1,
  green: -1,
  blue: -1,
};

function resetMinCubes() {
  minCubes.red = -1;
  minCubes.green = -1;
  minCubes.blue = -1;
}

function processLine(line) {
  resetMinCubes();

  let valid = true;
  const [idStr, dataStr] = line.split(": ");
  let id = parseInt(idStr.split(" ")[1]);
  const splitData = dataStr.split("; ");
  let color, num, set, splitSet;

  for (let i = 0; i < splitData.length; i++) {
    set = splitData[i];
    splitSet = set.split(", ");

    for (let j = 0; j < splitSet.length; j++) {
      [num, color] = splitSet[j].split(" ");
      num = parseInt(num);

      // Part 1
      if (num > cubes[color]) {
        valid = false;
      }

      // Part 2
      if (num > minCubes[color]) {
        minCubes[color] = num;
      }
    }
  }

  if (!valid) {
    id = -1;
  }

  const power = minCubes.red * minCubes.green * minCubes.blue;

  return [id, power];
}

function main() {
  const inputFileName =
    process.argv.includes("t") || process.argv.includes("-t")
      ? "input-test.txt"
      : "input.txt";

  const file = fs.readFileSync(inputFileName).toString();
  const lines = file.split("\n");
  let id, line, power;
  let sum = 0;
  let powerSum = 0;

  for (let i = 0; i < lines.length; i++) {
    line = lines[i];
    if (line === "") {
      continue;
    }

    [id, power] = processLine(line);
    if (id !== -1) {
      sum += id;
    }
    powerSum += power;
  }

  console.log("Sum of valid IDs: ", sum);
  console.log("Sum of powers: ", powerSum);
}

main();
