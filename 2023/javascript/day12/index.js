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
  /**LOG*/ logs.push(` >>>> Finding [${range[0]}, ${range[1]}] >>>>`);
  
  for (let cursor = range[0]; cursor < range[1] + 1; cursor++) {
    if (springs[cursor] === springStates.OPERATIONAL) {
      /**LOG*/ logs.push(`-!- range (${range}) is INVALID ::  currentStartIdx(${range[0]}) contains OPERATIONAL spring at (${cursor})`);
      /**LOG*/ logs.push(" >>>> ");
      return false;
    }
  }
  
  const nextIdx = range[1] + 1; 
  const nextCharIsValid = nextIdx === springs.length || springs[nextIdx] !== springStates.DAMAGED;
 
  return nextCharIsValid;
}

function walk(springs, range, groups, groupIdx, path) {
  if (groupIdx === groups.length - 1) {
    /**LOG*/ logs.push(` ### (BASE CASE) Last (#${groupIdx}) Group (${groups[groupIdx]})  and has range(${range})`);
    
    let str = Array(springs.length).fill(" ");
    let i, start, end;
    const fullPath = path.split("&");
    
    fullPath.forEach(rangeStr => {
      [ start, end ] = rangeStr.split("-");
      for (i = parseInt(start); i < parseInt(end) + 1; i++) {
        str[i] = "^";
      }
    });

    ///**LOG*/ console.log("full path:  ", fullPath);
    /**LOG*/ console.log("springs:    ", springs.join(""));
    /**LOG*/ console.log("possibility:", str.join(""));
    // /**LOG*/ logs.push(springs);
    // /**LOG*/ logs.push(str);
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
      //console.log("++ dispatching walk for group#", nextGroupIdx+1, "(", nextGroup, ") with range(", range, ") as group idx", nextGroupIdx);
      /**LOG*/ logs.push(`++ dispatching from #${groupIdx}->#${nextGroupIdx} walk(range=${range}, nextGroupIdx=${nextGroupIdx})`);
      possibilities += walk(springs, range, groups, nextGroupIdx, path + "&" + range.join("-"));
    }
      
    if (springs[start] === springStates.DAMAGED) {
      /**LOG*/ logs.push(`++ LAST START FOUND :: at range (${range})`);
      break;
    } else {
      /**LOG*/ logs.push(`++ last start not found spring(${springs[start]}) at start idx (${start})`);
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
  
  /**LOG*/ logs.push("\n---- found starting ranges ----\n");
  /**LOG*/ logs.push(
    "starting ranges: ",
    possibleStartingRanges.reduce((str, r) => str + `[${r[0]},${r[1]}],`, ""),
  );
  
  let possibilities = 0;
  possibleStartingRanges.forEach(range => {
    /**LOG*/ logs.push(" $ Starting Range:", range);
    possibilities += walk(springs, range, groups, 0, range.join("-"));
    /**LOG*/ logs.push("- - - - - - - - - - - - - - - - - - - - - ");
  });
  return possibilities;
}

var logs = {
  data: [],
  print() {
    this.data.forEach(l => console.log(l));
  },
  push(val) {
    this.data.push(val);
  },
  reset() {
    this.data = [];
  },
};

async function part1() {
  console.log("\n".repeat(10)+"* PART ONE *");
  let groups, springsStr, springs;
  let possibilities = 0;
  let result;
  let idx = 1;
  
  for await (const line of getLine()) {
    if (line === null) {
      break;
    }

    // Parse line into springs and groups
    [ springsStr, groups ] = line.split(" ");
    springs = springsStr.split("");
    groups = groups.split(",").map(numStr => parseInt(numStr));
    // Find possibilities for this line
    result = findPossibilityCount(springs, groups);
    possibilities += result;

    // /**LOG*/ logs.push(`Line "${idx++}" ( ${line+" )".padEnd(35, " ")} provided (${result}) possibilities`);
    /**LOG*/ console.log(`Line "${idx++}" ( ${line+" )".padEnd(35, " ")} provided (${result}) possibilities`);
   
    //logs.print();
    logs.reset();
  }
  
  return possibilities;
}

async function main() {
  const p1Result = await part1();
  /**LOG*/ console.log("\nPart 1 Result:", p1Result);
}

main();

