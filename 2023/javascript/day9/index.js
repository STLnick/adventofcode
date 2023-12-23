const fs = require("fs");

function getLines() {
  const filename = process.argv.includes("t") || process.argv.includes("-t")
    ? "input-test.txt"
    : "input.txt";
  return fs.readFileSync(__dirname + "/" + filename).toString().replaceAll("\r", "\n").split("\n").filter(Boolean);
}

function getFutureValue(history) {
  let workingSet = [ ...history ];
  let newSet;
  let lastVals = [ workingSet[workingSet.length - 1] ];
  let current, next;

  while (!workingSet.every(v => v === 0)) {
    newSet = [];

    for (let i = 0; i < workingSet.length - 1; i++) {
      current = workingSet[i];
      next = workingSet[i + 1];
      newSet.push(next - current);
    }

    workingSet = [ ...newSet ];
    lastVals.push(workingSet[workingSet.length - 1]);
  }
   
  const futureVal = lastVals.reduce((sum, v) => sum + v, 0);
  return futureVal;
}

function part1(historyList) {
  let futureValues = new Array(historyList.length).fill(null);
  let history;
  
  for (let i = 0; i < historyList.length; i++) {
    history = historyList[i];
    futureValues[i] = getFutureValue(history);
  }

  return futureValues.reduce((sum, v) => sum + v, 0);
}

function getPastValue(history) {
  let workingSet = [ ...history ];
  let newSet;
  let firstVals = [ workingSet[0] ];
  let current, next;

  while (!workingSet.every(v => v === 0)) {
    newSet = [];

    for (let i = 0; i < workingSet.length - 1; i++) {
      current = workingSet[i];
      next = workingSet[i + 1];
      newSet.push(next - current);
    }

    workingSet = [ ...newSet ];
    firstVals.unshift(workingSet[0]);
  }

  let nextVal;
  let pastVal = firstVals[0];
  for (let i = 1; i < firstVals.length; i++) {
    nextVal = firstVals[i];
    pastVal = nextVal - pastVal;
  }

  return pastVal;

}

function part2(historyList) {
  let pastValues = new Array(historyList.length).fill(null);
  let history;
  
  for (let i = 0; i < historyList.length; i++) {
    history = historyList[i];
    pastValues[i] = getPastValue(history);
  }

  return pastValues.reduce((sum, v) => sum + v, 0);
}

function main() {
  const historyList = getLines().map(s => s.split(" ").map(n => parseInt(n)));

  const p1Result = part1(historyList);
  console.log("P1 Result - future values sum:", p1Result);
  const p2Result = part2(historyList);
  console.log("P2 Result - past values sum:", p2Result);
}

main();

