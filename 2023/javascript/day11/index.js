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

function expandUniverse(lines, clearColumns, clearRows) {
  const copyStr = SPACE.repeat(lines[0].length);
  let inserted = 0;

  clearRows.forEach(rowIdx => {
    lines.splice(rowIdx + inserted, 0, String(copyStr).split("")); 
    inserted++;
  });

  inserted = 0;
  clearColumns.forEach(colIdx => {
    lines.forEach(line => {
      line.splice(colIdx + inserted, 0, SPACE);
    });
    inserted++;
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
  const pairDistances = findGalaxyPairPaths(lines);
  return pairDistances.reduce((sum, dist) => sum + dist, 0);
}

async function main() {
  const p1Result = await part1();
  /**LOG */ console.log("Part 1 Result:", p1Result);
}

main();
