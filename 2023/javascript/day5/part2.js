const fs = require("fs");

function createMap(lines, start) {
  let mappings = [];
  let finalMappings = [];

  for (let i = start; i < lines.length; i++) {
    if (lines[i] === "") {
      mappings.sort((a, b) => (a.src < b.src ? -1 : 1));
      
      let currentMapping, nextMapping;
      for (let j = 0; j < mappings.length; j++) {
        currentMapping = mappings[j];
        finalMappings.push(currentMapping);
        nextMapping = j + 1 < mappings.length ? mappings[j + 1] : null;

        if (nextMapping && currentMapping.srcEnd !== nextMapping.src - 1) {
          finalMappings.push({
            dest: currentMapping.srcEnd + 1,
            src: currentMapping.srcEnd + 1,
            srcEnd: nextMapping.src - 1,
            range: nextMapping.src - 1 - currentMapping.srcEnd,
            delta: 0,
          });
        }
      }
      
      finalMappings.sort((a, b) => (a.src < b.src ? -1 : 1));

      if (finalMappings[0].src !== 0) {
        finalMappings.unshift({
          dest: 0,
          src: 0,
          srcEnd: finalMappings[0].src - 1,
          range: finalMappings[0].src,
          delta: 0,
        });
      }

      const endSrc = finalMappings[finalMappings.length - 1].src
        + finalMappings[finalMappings.length - 1].range;
      
      finalMappings.push({
        dest: endSrc,
        src: endSrc,
        srcEnd: 5_000_000_000,
        range: 5_000_000_000 - endSrc,
        delta: 0,
      });

      finalMappings.sort((a, b) => (a.src < b.src ? -1 : 1));

      return [finalMappings, i];
    }

    const [dest, src, range] = lines[i].split(" ").map((x) => parseInt(x));
    mappings.push({
      dest,
      src,
      srcEnd: src + range - 1,
      range,
      delta: dest - src,
    });
  }
}

function splitRangeWithMappings(currentRange, mappings) {
  const start = currentRange[0] + 0;
  const end = currentRange[1] + 0;
  const mappingRange = mappings.find(m => valIsWithinRange(start, m));

  if (!mappingRange) {
    console.error("ERROR DUMP", { currentRange, mappings });
    throw new Error("unexpected null mappingRange");
  }

  if (valIsWithinRange(end, mappingRange)) {
    return [
      [ start + mappingRange.delta, end + mappingRange.delta ],
    ];
  }

  const nextStart = mappingRange.src + mappingRange.range;

  return [
    [ start + mappingRange.delta, nextStart - 1 + mappingRange.delta ],
    ...splitRangeWithMappings([nextStart, end], mappings),
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

  return seedRanges;
}

function getSeedRangeAsArray(range) {
  return new Array(range.start, range.start + range.range - 1);
}

function valIsWithinRange(val, range) {
  return val >= range.src && val <= range.src + range.range - 1;
}

function getLines() {
  const filename = process.argv.includes("-t") ? "input-test.txt" : "input.txt";
  return fs.readFileSync(filename).toString().split("\n");
}

function main(lines) {
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

  const seedRangeArrs = seedRanges.map(getSeedRangeAsArray);
  let currentRanges = [...seedRangeArrs];
  let mappings;

  for (let j = 0; j < Object.keys(masterMap).length; j++) {
    mappings = masterMap[j];
    const copiedRanges = JSON.parse(JSON.stringify(currentRanges));
    currentRanges = [];

    for (let ri = 0; ri < copiedRanges.length; ri++) {
      currentRanges.push(
        ...splitRangeWithMappings(copiedRanges[ri], mappings)
      );
    }
  
    // Re-sort the current ranges by start value
    currentRanges.sort((a, b) => (a[0] < b[0] ? -1 : 1));
  }

  return currentRanges[0][0];
}

const runs = 1;
const lines = getLines();
let val;
let results = {};

for (let m = 0; m < runs; m++) {
  val = main(lines);
  if (results[val]) {
    results[val] += 1;
  } else {
    results[val] = 1;
  }
}

console.log(`Ran (${runs}) times :: results\n${JSON.stringify(results, null, 2)}`);

