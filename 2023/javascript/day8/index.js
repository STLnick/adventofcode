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

function createNodeInObject(obj, dir) {
  const [ source, leftRightStr ] = dir.split(" = ");
  const [ left, right ] = leftRightStr.slice(1, leftRightStr.length - 1).split(", ");
  obj[source] = { L: left, R: right };
}

function part1(instructions, directions, start, goal) {
  let moves = 0;
  let current = start;
  let currentInstruction;
  let instructionIdx = 0;

  while (current !== goal) {
    currentInstruction = instructions[instructionIdx]; 
    current = directions[current][currentInstruction];
    moves++;
    instructionIdx++;

    if (instructionIdx === instructions.length) {
      instructionIdx = 0;
    }
  }

  return moves;
}

function main() {
  const [ instructions, ...directionStrs ] = getLines();
  const directions = {};
  directionStrs.forEach(ds => createNodeInObject(directions, ds));

  const part1Result = part1(instructions, directions, 'AAA', 'ZZZ');
  console.log({ part1Result });
}

main();

