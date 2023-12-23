const nums = [ 0, 1, 2, 3 ];

const NORTH = 0;
const EAST = 1;
const SOUTH = 2;
const WEST = 3;
const dirs = [ NORTH, EAST, SOUTH, WEST ];

const test2 = new Array(4).fill([]);
console.log({ test2 });
console.log(0, ":", test2[0]);
console.log(NORTH, ":", test2[NORTH]);

test2[0].push("A");

console.log(0, ":", test2[0]);
console.log(NORTH, ":", test2[NORTH]);

test2[0].push("B");

console.log(0, ":", test2[0]);
console.log(NORTH, ":", test2[NORTH]);
