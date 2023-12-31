const fs = require("fs");
const readline = require("readline");

const springStates = Object.freeze({
  DAMAGED: "#",
  OPERATIONAL: ".",
  UNKNOWN: "?",
});

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
  for (let cursor = range[0]; cursor < range[1] + 1; cursor++) {
    if (springs[cursor] === springStates.OPERATIONAL) {
      return false;
    }
  }
  
  const nextIdx = range[1] + 1; 
  const nextCharIsValid = nextIdx === springs.length || springs[nextIdx] !== springStates.DAMAGED;
 
  return nextCharIsValid;
}

function endHasDamaged(springs, start) {
  if (start >= springs.length) {
    return false;
  }

  for (let i = start; i < springs.length; i++) {
    if (springs[i] === springStates.DAMAGED) {
      return true;
    }
  }

  return false;
}

function walk(springs, range, groups, groupIdx, path) {
  const isLastGroup = groupIdx === groups.length - 1;
  
  if (isLastGroup) {
    return endHasDamaged(springs, range[1] + 2) ? 0 : 1;
  }

  const start = range[1] + 2; // Start with one space gap between last range
  const end = start + groups[groupIdx + 1] - 1;
  const newPath = `${path ? `${path}&`: ""}${range.join("-")}`;

  return calcPossibilities(springs, Array(start, end), groups, groupIdx + 1, newPath);
}

function calcPossibilities(springs, range, groups, groupIdx, path) {
  let start = range[0];
  let end = range[1];
  let possibilities = 0;
 
  for (; end < springs.length; start++, end++) {
    const range = Array(start, end);

    if (isPossibleRange(springs, range)) {
      const newPath = `${path ? `${path}&`: ""}${range.join("-")}`;
      possibilities += walk(springs, range, groups, groupIdx, newPath);
    }

    if (springs[start] === springStates.DAMAGED) {
      return possibilities;
    }
  }
  
  return possibilities;
}

async function run() {
  let groups, resultOne, resultTwo, springsStr, springs, temp;
  let possibilitiesOne = 0;
  let possibilitiesTwo = 0;
  
  for await (const line of getLine()) {
    if (line === null) {
      break;
    }

    [ springsStr, groups ] = line.split(" ");
    springs = springsStr.split("");
    groups = groups.split(",").map(numStr => parseInt(numStr));
    const startRange = Array(0, groups[0] - 1);

    /** Part One **/
    resultOne = calcPossibilities(springs, startRange, groups, 0, "");
    possibilitiesOne += resultOne;
    // console.log("ONE", { resultOne, springs, groups });

    /** Part Two **/
    if (springsStr.endsWith(springStates.DAMAGED)) {
      // Cannot use additional "?" just use resultOne in all multiplication spots
      possibilitiesTwo += Math.pow(resultOne, 5);
    } else {
      springs = `?${springsStr}`.split("");
      resultTwo = calcPossibilities(springs, startRange, groups, 0, "");
      // Swap the additional "?" to the other side and recalculate
      // -- when the first range cannot change at all?
      // -- if both ways change original possibility total - which should be used?
      //      - higher or lower? first way or second (adjusted) way?
      
      springs = `${springsStr}?`.split("");
      temp = calcPossibilities(springs, startRange, groups, 0, "");
      if (temp > resultTwo) {
        resultTwo = temp;
      }
      possibilitiesTwo += resultOne * Math.pow(resultTwo, 4);
      // console.log("TWO", { resultTwo, addedPoss: resultOne * Math.pow(resultTwo, 4), springs });
    }
  }
  
  return {
    p1Result: possibilitiesOne,
    p2Result: possibilitiesTwo,
  };
}

async function main() {
  const { p1Result, p2Result } = await run();
  /**LOG*/ console.log("\nPart 1 Result:", p1Result);
  /**LOG*/ console.log("Part 2 Result:", p2Result);
}

main();

