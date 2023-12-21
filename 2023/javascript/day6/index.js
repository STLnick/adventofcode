const fs = require("fs");

function getLines() {
  const filename = process.argv.includes("t") || process.argv.includes("-t")
    ? "input-test.txt"
    : "input.txt";
  return fs.readFileSync(__dirname + "/" + filename).toString().split("\n").filter(Boolean);
}

function getTimeDistanceMap() {
  const [time, distance] = getLines();
  const splitTime = time.split(": ")[1].trim().split(" ").filter(Boolean);
  const splitDistance = distance.split(": ")[1].trim().split(" ").filter(Boolean);

  let map = {};

  for (let i = 0; i < splitTime.length; i++) {
    map[+splitTime[i]] = +splitDistance[i];
  }

  return map;
}

function getMMTraveled(timeMs, msHeld) {
  if (msHeld === 0 || msHeld === timeMs) {
    return 0;
  }

  return (timeMs - msHeld) * msHeld;
}

function main() {
  const tdMap = getTimeDistanceMap();
  const times = Object.keys(tdMap);
  let possibilities = {};

  times.forEach(timeMs => {
    possibilities[timeMs] = 0;
    
    for (let held = 1; held < timeMs; held++) {
      if (getMMTraveled(timeMs, held) > tdMap[timeMs]) {
        possibilities[timeMs]++;
      }
    }
  });
  
  console.log(possibilities);
  const product = Object.values(possibilities).reduce((acc, num) => acc * num, 1);
  console.log(product);
}

main();

