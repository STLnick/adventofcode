const fs = require("fs");

function createMap(lines, start) {
  let mappings = [];

  for (let i = start; i < lines.length; i++) {
    if (lines[i] === "") {
      return [mappings, i];
    }

    const [dest, src, range] = lines[i].split(" ").map((x) => parseInt(x));
    mappings.push({ dest, src, range });
  }
}

const filename = process.argv.includes("-t") ? "input-test.txt" : "input.txt";
const lines = fs.readFileSync(filename).toString().split("\n");
let seeds;
let masterMap = {};
let mapIdx = 0;

for (let i = 0; i < lines.length; i++) {
  if (i === 0) {
    seeds = lines[i]
      .split(": ")[1]
      .split(" ")
      .map((x) => parseInt(x));
  } else if (lines[i] !== "") {
    const [newMap, resumeIdx] = createMap(lines, i + 1);
    masterMap[mapIdx++] = newMap;
    i = resumeIdx;
  }
}

let map;
let current;
let seed;
let seedMap = {};

for (let i = 0; i < seeds.length; i++) {
  current = seed = seeds[i];
  seedMap[seed] = [];

  for (let j = 0; j < Object.keys(masterMap).length; j++) {
    map = masterMap[j];

    for (let mi = 0; mi < map.length; mi++) {
      if (current >= map[mi].src && current <= map[mi].src + map[mi].range) {
        current = map[mi].dest + current - map[mi].src;
        break;
      }
    }

    seedMap[seed].push(current);
  }
}

const lowest = Object.keys(seedMap).reduce((low, key) => {
  const val = seedMap[key][seedMap[key].length - 1];
  if (val < low) {
    low = val;
  }
  return low;
}, Infinity);

console.log("Lowest Location Number: ", lowest);
