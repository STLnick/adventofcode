const fs = require("fs");
const readline = require("readline");

const springStates = {
  DAMAGED: "#",
  OPERATIONAL: ".",
  UNKNOWN: "?",
};

async function* getLine() {
  const filename = process.argv.includes("-t") ? "input-test.txt" : "input.txt";
  const fileStream = fs.createReadStream(filename);
  const reader = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });
 
  for await (const line of reader) {
    yield line;
  }
  yield null;
}

function isPossibleRange(springs, range) {
  for (let cursor = range[0], end = range[1]; cursor < end + 1; cursor++) {
    if (springs[cursor] === springStates.OPERATIONAL) {
      return false;
    }
  }
  
  const prevIdx = range[0] - 1; 
  const prevCharIsValid = prevIdx < 0 || springs[prevIdx] !== springStates.DAMAGED;
  
  if (!prevCharIsValid) {
    return false;
  }
  
  const nextIdx = range[1] + 1; 
  const nextCharIsValid = nextIdx >= springs.length || springs[nextIdx] !== springStates.DAMAGED;
  
  return nextCharIsValid;
}

function walkPossibilities(springs, range, groups, currentGroupIdx) {
  // if is a leaf node / possibility
  if (currentGroupIdx === groups.length - 1) {
    console.log("Should be LEAF NODE");
    console.log({ springs, range, groups, currentGroupIdx });
    console.log(groups[currentGroupIdx], ": range", range);

    return 1;
  }

  const nextGroupIdx = groups[currentGroupIdx + 1];
  const nextGroup = groups[nextGroupIdx];
  let start = range[1];

  if (currentGroupIdx !== -1) {
    start += 2; // Start with one space gap between last range
  }

  let end = start + nextGroup - 1;
  
  if (end >= springs.length) {
    return 0;
  }
  
  let possibilities = 0;

  while (end < springs.length) {
    const range = new Array(start, end);
    if (isPossibleRange(springs, range)) {
      possibilities += walkPossibilities(springs, range, groups, nextGroupIdx);
    }

    start++;
    end++;
  }

  return possibilities;
}

function findRange(length, offset) {
  return new Array(offset, offset + length - 1);
}

function findPossibilites(springs, groups) {
  console.log("findPossibilites() :: START", { springs, groups });
  let possibilities = 0;

  for (let i = 0; i < groups.length; i++) {
    possibilities += walkPossibilities(springs, new Array(0, 0), groups, -1);
    console.log("* possibilities", possibilities);
  }
}

async function part1() {
  let groups, springsStr, springs;

  for await (const line of getLine()) {
    if (line === null) {
      break;
    }

    [ springsStr, groups ] = line.split(" ");
    springs = springsStr.split("");
    groups = groups.split(",").map(numStr => parseInt(numStr));

    findPossibilites(springs, groups);
  }
}

async function main() {
  const p1Result = await part1();
  /**LOG*/ console.log("\nPart 1 Result:", p1Result);
}

main();

