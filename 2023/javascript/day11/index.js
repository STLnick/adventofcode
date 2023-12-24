const fs = require("fs");
const readline = require("readline");

const GALAXY = "#";
const SPACE = ".";

async function* processLineByLine() {
  const filename = process.argv.includes("t") || process.argv.includes("-t")
    ? "input-test.txt"
    : "input.txt";
  const fileStream = fs.createReadStream(filename);
  const reader = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });

  for await (const line of reader) {
    yield line.split("");
  }
  yield null;
}

async function parseData() {
  let clearColumns = [];
  let clearRows = [];
  let lines = [];
  let y = 0;

  for await (const line of processLineByLine()) {
    if (line === null) {
      break;
    }
   
    lines.push(line);

    if (line.every(c => c === SPACE)) {
      clearRows.push(y);
    }

    y++;
  }

  for (let x = 0; x < lines[0].length; x++) {
    if (lines.every(l => l[x] === SPACE)) {
      clearColumns.push(x);
    }
  }

  return { clearColumns, clearRows, lines };
}

function findGalaxies(lines) {
  let galaxies = [];
  let id = 1;

  lines.forEach((line, y) => {
    line.forEach((char, x) => {
      if (char === GALAXY) {
        galaxies.push({ id, coords: [y, x] });
        id++;
      }
    });
  });

  return galaxies;
}

function expandUniverse(lines, clearColumns, clearRows, expandFactor = 2) {
  const copyRow = SPACE.repeat(lines[0].length);
  let inserted = 0;

  clearRows.forEach(rowIdx => {
    const newRows = new Array(expandFactor - 1).fill(String(copyRow).split(""));
    lines.splice(rowIdx + inserted, 0, ...newRows); 
    inserted += expandFactor - 1;
  });

  inserted = 0;
  clearColumns.forEach(colIdx => {
    lines.forEach(line => {
      const newCols = new Array(expandFactor - 1).fill(SPACE);
      line.splice(colIdx + inserted, 0, ...newCols);
    });
    inserted += expandFactor - 1;
  });
}

function getNumberOfPossiblePairs(numOfNodes) {
  let num = 0;
  for (let add = numOfNodes - 1; add > 0; add--) {
    num += add;
  }
  return num;
}

function getDistanceBetweenGalaxies([y1, x1], [y2, x2]) {
  return Math.abs(y2 - y1) + Math.abs(x2 - x1);
}

function findGalaxyPairPaths(lines) {
  const galaxies = findGalaxies(lines);
  const numOfPossiblePairs = getNumberOfPossiblePairs(galaxies.length);
  let distances = new Array(numOfPossiblePairs);
  let galaxy, pairGalaxy;
  let distIdx = 0;

  for (let i = 0; i < galaxies.length - 1; i++) {
    galaxy = galaxies[i];
 
    for (let j = i + 1; j < galaxies.length; j++) {
      pairGalaxy = galaxies[j];
      distances[distIdx] = getDistanceBetweenGalaxies(galaxy.coords, pairGalaxy.coords),
      distIdx++;
    }
  }
  
  return distances;
}

async function part1() {
  let { clearColumns, clearRows, lines } = await parseData();
  expandUniverse(lines, clearColumns, clearRows);
  const doubledUniSum = findGalaxyPairPaths(lines).reduce((sum, dist) => sum + dist, 0);
  return doubledUniSum;
}

async function part2(factor) {
  let { clearColumns, clearRows, lines } = await parseData();
  const baseSum = findGalaxyPairPaths(lines).reduce((sum, dist) => sum + dist, 0);
  expandUniverse(lines, clearColumns, clearRows);
  const doubledUniSum = findGalaxyPairPaths(lines).reduce((sum, dist) => sum + dist, 0);
  const diff = doubledUniSum - baseSum;

  return baseSum + diff * (factor - 1);
}

async function main() {
  /**LOG */ console.log("-- Day 11 --");
  const p1Result = await part1();
  /**LOG */ console.log("Part 1 Result:", p1Result);
  const p2Result = await part2(1_000_000);
  /**LOG */ console.log("Part 2 Result:", p2Result);
}

main();

