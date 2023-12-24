const fs = require("fs");
const readline = require("readline");

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

async function findClearRowsCols() {
  let clearColumns = [];
  let clearRows = [];
  let lines = [];
  let y = 0;

  for await (const line of processLineByLine()) {
    if (line === null) {
      continue;
    }
   
    lines.push(line);

    if (line.every(c => c === ".")) {
      clearRows.push(y);
    }

    y++;
  }

  for (let x = 0; x < lines[0].length; x++) {
    if (lines.every(l => l[x] === ".")) {
      clearColumns.push(x);
    }
  }

  lines.forEach(l => console.log(l.join("")));

  return { clearColumns, clearRows, lines };
}

function expandUniverse(lines, clearColumns, clearRows) {
  const copyStr = String(".").repeat(lines[0].length);
  let inserted = 0;

  clearRows.forEach(rowIdx => {
    lines.splice(rowIdx + inserted, 0, String(copyStr).split("")); 
    inserted++;
  });

  inserted = 0;
  clearColumns.forEach(colIdx => {
    lines.forEach(line => {
      line.splice(colIdx + inserted, 0, ".");
    });
    inserted++;
  });
}

async function part1() {
  let { clearColumns, clearRows, lines } = await findClearRowsCols();
  console.log({ clearRows, clearColumns });

  lines = expandUniverse(lines, clearColumns, clearRows);

  return 'IDK';
}

async function main() {
  const p1Result = await part1();
  console.log("Part One Result:", p1Result);
}

main();
