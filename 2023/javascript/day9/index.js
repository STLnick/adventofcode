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

function part1() {
  const historyList = getLines().map(s => s.split(" ").map(n => parseInt(n)));
  let futureValues = new Array(historyList.length).fill(null);
  let history;
  
  for (let i = 0; i < historyList.length; i++) {
    history = historyList[i];
    futureValues[i] = getFutureValue(history);
  }

  return futureValues.reduce((sum, v) => sum + v, 0);
}

function main() {
  const p1Result = part1();
  console.log("P1 Result - future values sum:", p1Result);
}

main();

