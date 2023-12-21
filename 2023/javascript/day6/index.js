const fs = require("fs");

function getLines() {
  const filename = process.argv.includes("t") || process.argv.includes("-t")
    ? "input-test.txt"
    : "input.txt";
  return fs.readFileSync(__dirname + "/" + filename).toString().split("\n").filter(Boolean);
}

function getTimeDistanceMap(joinAsOneNum = false) {
  const [time, distance] = getLines();
  let splitTime = time.split(": ")[1].trim().split(" ").filter(Boolean);
  let splitDistance = distance.split(": ")[1].trim().split(" ").filter(Boolean);

  if (joinAsOneNum) {
    splitTime = [ splitTime.join("") ];
    splitDistance = [ splitDistance.join("") ];
  }

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

function getPossibilities(tdMap) {
  const times = Object.keys(tdMap);
  let possibilities = {};
  let reachedPossibles;

  times.forEach(timeMs => {
    possibilities[timeMs] = 0;
    reachedPossibles = false;
    
    for (let held = 1; held < timeMs; held++) {
      if (getMMTraveled(timeMs, held) > tdMap[timeMs]) {
        if (!reachedPossibles) reachedPossibles = true
        possibilities[timeMs]++;
      } else if (reachedPossibles) {
        break;
      }
    }
  });

  return possibilities;
}

function part1() {
  const tdMap = getTimeDistanceMap();
  
  let possibilities = getPossibilities(tdMap);
  console.log("P1:", possibilities);
  
  const product = Object.values(possibilities).reduce((acc, num) => acc * num, 1);
  console.log("P1:", product);
}

function part2() {
  const tdMap = getTimeDistanceMap(true);
  
  let possibilities = getPossibilities(tdMap);
  console.log("P2", possibilities);
  
  const product = Object.values(possibilities).reduce((acc, num) => acc * num, 1);
  console.log("P2", product);
}

function part2_1() {
  const tdMap = getTimeDistanceMap(true);
  const timeMs = Object.keys(tdMap)[0];
  const recordDist = tdMap[timeMs];
  const msHeld = Math.floor(+timeMs / 2);
  const root = Math.floor(Math.sqrt(+timeMs));
  let start = msHeld;
  let end = msHeld;

  //work begin leftwards until edge
  while (start > 0 && getMMTraveled(+timeMs, start) > recordDist) {
    start -= root;
  }
  start += root;

  while (start > 0 && getMMTraveled(+timeMs, start) > recordDist) {
    start -= 1;
  }
  start += 1;
  
  //work end rightwards until edge
  while (end < timeMs && getMMTraveled(+timeMs, end) > recordDist) {
    end += root;
  }
  end -= root;

  while (end < timeMs && getMMTraveled(+timeMs, end) > recordDist) {
    end += 1;
  }
  end -= 1;

  const possiblities = end - start + 1;
  console.log("P2_1", possiblities);
}

console.time("p1")
part1();
console.timeEnd("p1")

console.time("p2")
part2();
console.timeEnd("p2")

console.time("p2_1")
part2_1();
console.timeEnd("p2_1")

