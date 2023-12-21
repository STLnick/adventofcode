const fs = require("fs");
let LOG = true;

function createMap(lines, start) {
  let mappings = [];

  for (let i = start; i < lines.length; i++) {
    if (lines[i] === "") {
      mappings.sort((a, b) => a.src < b.src ? -1 : 1);

      // 0 to X range?
      if (mappings[0].src !== 0) {
        mappings.unshift({
          dest: 0,
          src: 0,
          srcEnd: mappings[0].src - 1,
          range: mappings[0].src,
          delta: 0,
        });
      }

      // N to Infinity range
      mappings.push({
        dest: mappings[mappings.length - 1].src + mappings[mappings.length - 1].range,
        src: mappings[mappings.length - 1].src + mappings[mappings.length - 1].range,
        srcEnd: Infinity,
        range: Infinity,
        delta: 0,
      });
      
      // Fill in gaps ranges
      let rangeOne, rangeTwo;
      for (let g = 0; g < mappings.length - 3; g++) {
        // if i had 2 mappings (-2 for those added above) I have one gap to check
        rangeOne = mappings[g+1]; // [1] src 50 range 48 - up to 97
        rangeTwo = mappings[g+2]; // [2] src 98 - NO GAP
        console.log({ rangeOne, rangeTwo });

        if (rangeOne.src + rangeOne.range === rangeTwo.src) {
          console.log("NO GAP");
        } else {
          console.log("Create range for 1-to-1 GAP");
        }
      }

      return [mappings, i];
    }

    const [dest, src, range] = lines[i].split(" ").map((x) => parseInt(x));
    mappings.push({ dest, src, srcEnd: src + range - 1, range, delta: dest - src });
  }
}

function splitRangeWithMappings(currentRange, mappings) {
  console.log("splitRangeWithMappings() :: currentRange", currentRange);
  let mappingRangeIdx = mappings.findIndex(m => valIsWithinRange(currentRange[0], m));
  let mappingRange = mappingRangeIdx !== -1 ? mappings[mappingRangeIdx] : null;

  if (valIsWithinRange(currentRange[1], mappingRange)) {
    return [
      [ currentRange[0] + mappingRange.delta, currentRange[1] + mappingRange.delta ]
    ];
  }

  return [
    [ currentRange[0] + 0, mappingRange.src + mappingRange.range - 1 ],
    ...splitRangeWithMappings(
      [ mappingRange.src + mappingRange.range, currentRange[1] + 0 ],
      mappings,
    ),
  ];
}

function getSeeds(line) {
  let nums = line
    .split(": ")[1]
    .split(" ")
    .map((x) => parseInt(x));
  let start;
  let range;
  let seedRanges = [];

  for (let i = 0; i < nums.length; i = i + 2) {
    start = nums[i];
    range = nums[i + 1];
    seedRanges.push({ start, range });
  }

  if (LOG) console.log("seed ranges: ", seedRanges);
  return seedRanges;
}
 
function getSeedRangeAsArray(range) {
  return new Array(range.start, range.start + range.range - 1);
}

function valIsWithinRange(val, range) {
  return val >= range.src && val <= range.src + range.range - 1;
}

const filename = process.argv.includes("-t") ? "input-test.txt" : "input.txt";
const lines = fs.readFileSync(filename).toString().split("\n");
const seedRanges = getSeeds(lines[0]);
let masterMap = {};
let mapIdx = 0;

for (let i = 1; i < lines.length; i++) {
  if (lines[i] !== "") {
    const [newMap, resumeIdx] = createMap(lines, i + 1);
    masterMap[mapIdx++] = newMap;
    i = resumeIdx;
  }
}

let current;
let lowest = Infinity;

if (LOG) {
  console.log("----------------");
  console.log("Master Map:");
  console.log(masterMap);
  console.log("----------------");
}

const seedRangeArrs = seedRanges.map(getSeedRangeAsArray);
if (LOG) console.log("seedRangeArrs", seedRangeArrs);

// [ [79, 92], [55, 67] ]
let currentRanges = [ ...seedRangeArrs ];
let mappings, mappingRange, mappingRangeIdx, workingRange;

for (let j = 0; j < Object.keys(masterMap).length; j++) {
  // Go through all steps splitting ranges as needed

  mappings = masterMap[j];
  console.log("Mappings:", mappings);

  // Re-map all of our current ranges according to this Step
  const copiedRanges = JSON.parse(JSON.stringify(currentRanges));
  currentRanges = []; // Empty ranges to push in results from each step

  for (let ri = 0; ri < copiedRanges.length; ri++) {
    workingRange = copiedRanges[ri];
      
    //TESTING NEW METHOD
    console.log("$$$$$$$$ provided:", workingRange);
    console.log("$$$$$$$$ result:", splitRangeWithMappings(workingRange, mappings));

    // Where our `start` (current sub-range) begins
    mappingRangeIdx = mappings.findIndex(m => valIsWithinRange(workingRange[0], m));
    mappingRange = mappingRangeIdx !== -1 ? mappings[mappingRangeIdx] : null;

    if (!mappingRange) {
      console.log(" -- No Range Found :: 1-to-1 mapping");
      // IS this workingRange entirely in a 1-to-1 range?

      // Where is the end of the sub-range established by workingRange[0]?
      // AKA where do we hit a different mapping Range?
    } else {
      console.log(" -- Found Range:", mappingRange);

     
      // IS the whole range in the same mapping range?
      if (valIsWithinRange(workingRange[1], mappingRange)) {
        console.log(" -- -- The entire workingRange(", workingRange, ") is within range: ", mappingRange);
        
        workingRange[0] += mappingRange.delta;
        workingRange[1] += mappingRange.delta;
        
        console.log(" -- -- ++Modified workingRange: ", workingRange);
        currentRanges.push(workingRange);
      } else {
        console.log(" -- -- workingRange: ", workingRange);
        console.log(" -- -- The workingRange must be split...");

        // Where is the end of the sub-range established by workingRange[0]?
        // [Start - X], ..., [N+1, End]
        //    [57-69], [81-94]
        //    /    \       \
        //   /      \       \
        // [57-60],[61-69],[81-94]

        // Handle part of subrange that's within current mappingRange
        const nextStart = mappingRange.src + mappingRange.range;
        const newRange = new Array(workingRange[0] + 0, nextStart - 1);
        newRange[0] += mappingRange.delta;
        newRange[1] += mappingRange.delta;
        // newRange = [57-60]
        currentRanges.push(newRange);

        mappingRangeIdx = mappings.findIndex(m => valIsWithinRange(nextStart, m));
        mappingRange = mappingRangeIdx !== -1 ? mappings[mappingRangeIdx] : null;
        
        if (!mappingRange) {
          throw new Error("mapping range not found");
        }

        // Find next subrange section
        const newRange2 = new Array(nextStart, workingRange[1] + 0);
        if (valIsWithinRange(newRange2[1], mappingRange)) {
          newRange2[0] += mappingRange.delta;
          newRange2[1] += mappingRange.delta;
        } else {
          console.log("!!!need to split again!!!");
          console.log("vals and such: ", { workingRange, newRange, newRange2, endValue: newRange2[1], mappingRange, mappings });
          throw new Error("pls no");
        }

        currentRanges.push(newRange2);
      }
    }
  }

  // Re-sort the current ranges by start value
  currentRanges.sort((a, b) => a[0] < b[0] ? -1 : 1);
  console.log("Current Ranges:", currentRanges);
}

console.log("Lowest Location Number [first range, first val]: ", currentRanges[0][0]);
