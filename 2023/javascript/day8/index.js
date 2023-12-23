const fs = require("fs");

function getLines() {
  let filename = "input.txt";
  if (process.argv.includes("t") || process.argv.includes("-t")) {
    filename = "input-test.txt";
  } else if (process.argv.includes("t2") || process.argv.includes("-t2")) {
    filename = "input-test-p2.txt";
  }

  return fs.readFileSync(__dirname + "/" + filename).toString().split("\n").filter(Boolean);
}

function createDirectionNodes(directionStrs) {
  const directions = {};
  directionStrs.forEach(ds => {
    const [ source, leftRightStr ] = ds.split(" = ");
    const [ left, right ] = leftRightStr.slice(1, leftRightStr.length - 1).split(", ");
    directions[source] = { L: left, R: right };
  });
  return directions;
}

function part1(instructions, directions, start, goal) {
  let moves = 0;
  let current = start;
  let currentInstruction;
  let instructionIdx = 0;

  while (current !== goal) {
    currentInstruction = instructions[instructionIdx]; 
    current = directions[current][currentInstruction];
    moves++;
    instructionIdx++;

    if (instructionIdx === instructions.length) {
      instructionIdx = 0;
    }
  }

  return moves;
}

function isPrime(num) {
  for (let i = 2, s = Math.sqrt(num); i <= s; i++) {
    if (num % i === 0) {
      return false;
    }
  }

  return num > 1;
}

let factors = [];
for (let f = 2; f < 1000; f++) {
  if (isPrime(f)) {
    factors.push(f);
  }
}

function findSmallestFactor(num) {
  return factors.find(f => num % f === 0);
}

function getPrimeFactors(num) {
  let factors = [];
  let smallestFactor;
  let quotient = num;
  
  while (quotient > 1) {
    smallestFactor = findSmallestFactor(quotient);
    factors.push(smallestFactor);
    quotient = quotient / smallestFactor;
  }

  return factors;
} 

function getStepCountsUntilZ(instructions, directions) {
  const nodes = Object.keys(directions).filter(node => node.endsWith("A"));
  let stepCounts = new Array(nodes.length).fill(-1);
  let currentInstruction;
  let instructionIdx = 0;
  let steps = 0;
  let stepIdx = 0;

  while (nodes.length) {
    currentInstruction = instructions[instructionIdx]; 
    
    for (let i = 0; i < nodes.length; i++) {
      nodes[i] = directions[nodes[i]][currentInstruction];
    }
    
    steps++;
    
    while (nodes.some(n => n.endsWith("Z"))) {
      const index = nodes.findIndex(n => n.endsWith("Z"));
      nodes.splice(index, 1);

      stepCounts[stepIdx] = steps;
      stepIdx++;
    }

    instructionIdx++;
    if (instructionIdx === instructions.length) {
      instructionIdx = 0;
    }
  }

  return stepCounts;
}

function removeFactorFromLists(lists, factor) {
  let idx = -1;
  lists.forEach(list => {
    idx = list.findIndex(f => f === factor);
    if (idx !== -1) {
      list.splice(idx, 1);
    }
  });
}

function part2(instructions, directions) {
  const stepCounts = getStepCountsUntilZ(instructions, directions);
  const stepCountPrimeFactors = stepCounts.reduce((arr, sc) => {
    arr.push(getPrimeFactors(sc));
    return arr;
  }, []);

  let primeFactors = [];
  let val;

  stepCountPrimeFactors.forEach(factorList => {
    while (factorList.length) {
      val = factorList[0];
      primeFactors.push(val);
      removeFactorFromLists(stepCountPrimeFactors, val);
    }
  });

  const lowestCommonMultiple = primeFactors.reduce((acc, f) => acc * f, 1);

  return lowestCommonMultiple;
}

function main() {
  const [ instructions, ...directionStrs ] = getLines();
  const directions = createDirectionNodes(directionStrs);
  
  const part1Result = part1(instructions, directions, 'AAA', 'ZZZ');
  console.log({ part1Result });
  
  const part2Result = part2(instructions, directions);
  console.log({ part2Result });
}

main();

