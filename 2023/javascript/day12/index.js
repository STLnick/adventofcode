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

async function part1() {
  let groups, springs;

  for await (const line of getLine()) {
    if (line === null) {
      break;
    }

    [ springs, groups ] = line.split(" ");
    springs = springs.split("");
    groups = groups.split(",");

    /**LOG*/ console.log({ line, springs, groups });
  }
}

async function main() {
  const p1Result = await part1();
  /**LOG*/ console.log("Part 1 Result:", p1Result);
}

main();

