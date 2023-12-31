const fs = require("fs");
const readline = require("readline");

const logs = {
  data: [],
  disabled: false,
  disableLogs() {
    this.disabled = true;
  },
  print() {
    if (this.disabled) return;
    this.data.forEach(l => console.log(l));
  },
  push(val) {
    if (this.disabled) return;
    this.data.push(val);
  },
  reset() {
    if (this.disabled) return;
    this.data = [];
  },
};
const springStates = Object.freeze({
  DAMAGED: "#",
  OPERATIONAL: ".",
  UNKNOWN: "?",
});
const walkEvents = Object.freeze({
  DISPATCH: 'dispatch',
  POSSIBILITY: 'possibility',
  RUN: 'run',
});

function logWalk({ eventType, springs, range, groups, groupIdx, path, callback }) {
  switch (eventType) {
    case walkEvents.DISPATCH:
      logs.push(`++ dispatching from #${groupIdx}->#${groupIdx + 1} walk(range=${range}, nextGroupIdx=${groupIdx + 1})`);
      break;
    case walkEvents.POSSIBILITY:
      let str = Array(springs.length).fill(" ");
      let i, start, end;
      path.split("&").forEach(rangeStr => {
        [ start, end ] = rangeStr.split("-");
        for (i = parseInt(start); i < parseInt(end) + 1; i++) {
          str[i] = "^";
        }
      });

      logs.push(
        ` ### (BASE CASE) Last (#${groupIdx}) Group (${groups[groupIdx]}) and has range(${range})`
        + "\n\tsprings:" + springs.join("")
        + "\n\tpssblty:" + str.join("")
      );
      break;
    case walkEvents.RUN:
      logs.push(" $ Starting Range:", range);
      callback();
      logs.push("- - - - - - - - - - - - - - - - - - - - - ");
      break;
  }
}

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
  logs.push(` >>>> Finding [${range[0]}, ${range[1]}] >>>>`);
  
  for (let cursor = range[0]; cursor < range[1] + 1; cursor++) {
    if (springs[cursor] === springStates.OPERATIONAL) {
      logs.push(`-!- range (${range}) is INVALID ::  currentStartIdx(${range[0]}) contains OPERATIONAL spring at (${cursor})`);
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
  if (groupIdx === groups.length - 1) {
    if (endHasDamaged(springs, range[1] + 2)) {
      return 0;
    }

    logWalk({ eventType: walkEvents.POSSIBILITY, springs, range, groups, groupIdx, path });
    return 1;
  }

  const nextGroupIdx = groupIdx + 1;
  const nextGroup = groups[nextGroupIdx];
  let start = range[1] + 2; // Start with one space gap between last range
  let end = start + nextGroup - 1;
  let possibilities = 0;

  for (; end < springs.length; start++, end++) {
    const range = new Array(start, end);

    if (isPossibleRange(springs, range)) {
      logWalk({ eventType: walkEvents.DISPATCH, range, groupIdx });
      possibilities += walk(springs, range, groups, nextGroupIdx, path + "&" + range.join("-"));
    }
      
    if (springs[start] === springStates.DAMAGED) {
      break;
    }
  }
        
  return possibilities;
}

function findPossibilityCount(springs, groups) {
  const possibleStartingRanges = [];
  const startingGroup = groups[0];
 
  for (let i = 0, j = startingGroup - 1; j < springs.length; i++, j++) {
    const range = Array(i, j);
    if (isPossibleRange(springs, range)) {
      possibleStartingRanges.push(range);
    }

    if (springs[i] === springStates.DAMAGED) {
      break;
    }
  }
  
  logs.push(
    "starting ranges: ",
    possibleStartingRanges.reduce((str, r) => str + `[${r[0]},${r[1]}],`, ""),
  );
  
  let possibilities = 0;
  possibleStartingRanges.forEach(range => {
    logWalk({
      eventType: walkEvents.RUN,
      range,
      callback: () => possibilities += walk(springs, range, groups, 0, range.join("-")),
    });
  });
  return possibilities;
}

async function part1() {
  logs.push("\n".repeat(10) + "* PART ONE *");
  let groups, springsStr, springs;
  let possibilities = 0;
  let result;
  let idx = 1;
  
  for await (const line of getLine()) {
    if (line === null) {
      break;
    }

    [ springsStr, groups ] = line.split(" ");
    // TESTING PART TWO
    springsStr = `${springsStr}?${springsStr}`
    groups = `${groups},${groups}`
    // - - - - - - - - - - - - - - - - - - - - - - 
    springs = springsStr.split("");
    groups = groups.split(",").map(numStr => parseInt(numStr));

    result = findPossibilityCount(springs, groups); // Now the result of one of four sections in "unfolded records"
    console.log("possibilites for a two-section:", result);
    result = Math.pow(result * 4, 2);
    console.log("possibilites for a FULL unfolded section:", result);

    possibilities += result;

    logs.push(`Line "${idx++}" ( ${springsStr + groups +" )".padEnd(35, " ")} provided (${result}) possibilities`);
    logs.print();
    logs.reset();
  }
  
  return possibilities;
}

async function main() {
  if (process.argv.includes("-n")) {
    logs.disableLogs();
  }
  const p1Result = await part1();
  /**LOG*/ console.log("\nPart 1 Result:", p1Result);
}

main();

