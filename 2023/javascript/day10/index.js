const fs = require("fs");

function getLines() {
  let filename = "input.txt";
  if (process.argv.includes("t") || process.argv.includes("-t")) {
    filename = "input-test.txt";
  } else if (process.argv.includes("t2") || process.argv.includes("-t2")) {
    filename = "input-test2.txt";
  }
  return fs.readFileSync(__dirname + "/" + filename).toString().split("\n").filter(Boolean);
}

const NORTH = 0;
const EAST = 1;
const SOUTH = 2;
const WEST = 3;
const pipeChart = {
  "|": [NORTH, SOUTH],
  "-": [EAST, WEST],
  "L": [NORTH, EAST],
  "J": [NORTH, WEST],
  "7": [SOUTH, WEST],
  "F": [EAST, SOUTH],
  ".": [null, null],
};
let dir1, dir2;
const pipesByDirection = Object.keys(pipeChart).reduce((acc, pipe) => {
  dir1 = pipeChart[pipe][0];
  dir2 = pipeChart[pipe][1];
  if (!acc[dir1]) {
    acc[dir1] = [ pipe ];
  } else {
    acc[dir1].push(pipe);
  }
  if (!acc[dir2]) {
    acc[dir2] = [ pipe ];
  } else {
    acc[dir2].push(pipe);
  }
  return acc;
}, {});
const checkCoordinates = [
  { coords: [1, 0], eligible: pipesByDirection[NORTH] },
  { coords: [0, -1], eligible: pipesByDirection[EAST] },
  { coords: [-1, 0], eligible: pipesByDirection[SOUTH] },
  { coords: [0, 1], eligible: pipesByDirection[WEST], },
];

function getStartCoords(lines) {
  const rows = lines.length;
  const cols = lines[0].length;
  let char, row;
  
  for (let y = 0; y < rows; y++) {
    row = lines[y];
    for (let x = 0; x < cols; x++) {
      char = row[x];
      if (char === "S") {
        return new Array(y, x);
      }
    }
  }
}

function getStartingPaths(startCoords, lines) {
  let path1, path2;
  let y, x;

  for (let dir = 0; dir < checkCoordinates.length; dir++) {
    const { coords, eligible } = checkCoordinates[dir];
    y = startCoords[0] + coords[0];
    x = startCoords[1] + coords[1];

    if (eligible.includes(lines[y][x])) {
      if (!path1) {
        path1 = { char: lines[y][x], coords: [ y, x ], source: dir };
      } else {
        path2 = { char: lines[y][x], coords: [ y, x ], source: dir };
        break;
      }
    }
  }

  return new Array(path1, path2);
}

function onSameSquare(coordsA, coordsB) {
  return coordsA[0] === coordsB[0] && coordsA[1] === coordsB[1];
}

function determineSource(dir) {
  switch (dir) {
    case NORTH:
      return SOUTH;
    case EAST:
      return WEST;
    case SOUTH:
      return NORTH;
    case WEST:
      return EAST;
  }
}

function advancePath(path, lines) {
  const dest = pipeChart[path.char][0] !== path.source
    ? pipeChart[path.char][0]
    : pipeChart[path.char][1];

  switch (dest) {
    case NORTH:
      path.coords[0] -= 1;
      break;
    case EAST:
      path.coords[1] += 1;
      break;
    case SOUTH:
      path.coords[0] += 1;
      break;
    case WEST:
      path.coords[1] -= 1;
      break;
  }
  
  path.char = lines[path.coords[0]][path.coords[1]];
  path.source = determineSource(dest);
}

function part1() {
  const lines = getLines();
  const startRowCol = getStartCoords(lines);
  const [ path1, path2 ] = getStartingPaths(startRowCol, lines);
  let steps = 1;

  while (!onSameSquare(path1.coords, path2.coords)) {
    steps++;
    advancePath(path1, lines);
    advancePath(path2, lines);
  }

  return steps;
}

function main() {
  console.log("-- Day 10 --");

  const p1Result = part1();
  console.log("Part One Result: ", p1Result);
}

main();

