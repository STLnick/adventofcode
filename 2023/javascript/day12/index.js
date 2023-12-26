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

function checkhighestOfGroups(groups, num) {
  const highest = groups.reduce((hi, n) => hi < n ? n : hi, 0);
  return highest === num;
}

function processSpringsStr(springs, groups) {
  console.log("\n- - - - - - - - - - - - - - - - - - - - -");
  console.log("processSpringsStr() :: ", { springs: springs.join(""), groups });

  let checkIdx, end, group, start;
  let nextStart = 0;
  let prevEnd = 0;
  let possible = false;

  for (let g = 0; g < groups.length; g++) {
    group = parseInt(groups[g]);
    start = parseInt(nextStart);
    end = start + group;
    possible = false;
    checkIdx = 0;

    // start to walk string to find a possibility
    // find a X length str of "#" or "?" or a combo
    console.log("processSpringsStr() :: group", group, ":", { start, end });
    
    while (!possible) {
      possible = true;

      for (checkIdx = start; checkIdx < end; checkIdx++) {
        if (springs[checkIdx] === springStates.OPERATIONAL) {
          possible = false;
          break;
        }
      }
   
      if (!possible) {
        start++;
        end++;
      }
    }

    console.log("* Possible section for(", group, "):(", start, ",", end, "):", springs.slice(start, end));

    prevEnd = end;
    nextStart = end + 1;
  }
}

async function part1() {
  let groups, springsStr, springs;

  let logged = false;
  for await (const line of getLine()) {
    if (line === null) {
      break;
    }

    [ springsStr, groups ] = line.split(" ");
    springs = springsStr.split("");
    groups = groups.split(",");

    // processSpringsStr(springs, groups);
    if (!logged) {
      processSpringsStr(springs, groups);
      logged = true;
    }

    // /**LOG*/ console.log({ line, springs, groups });
  }
}

async function main() {
  const p1Result = await part1();
  /**LOG*/ console.log("Part 1 Result:", p1Result);
}

main();

