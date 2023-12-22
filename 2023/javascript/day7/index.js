const fs = require("fs");

function getLines() {
  const filename = process.argv.includes("t") || process.argv.includes("-t")
    ? "input-test.txt"
    : "input.txt";
  return fs.readFileSync(__dirname + "/" + filename).toString().split("\n").filter(Boolean);
}

const HAND_LENGTH = 5;
const cardRanks = {
  '2': 0,
  '3': 1,
  '4': 2,
  '5': 3,
  '6': 4,
  '7': 5,
  '8': 6,
  '9': 7,
  'T': 8,
  'J': 9,
  'Q': 10,
  'K': 11,
  'A': 12,
};
const cardRanks2 = {
  'J': 0,
  '2': 1,
  '3': 2,
  '4': 3,
  '5': 4,
  '6': 5,
  '7': 6,
  '8': 7,
  '9': 8,
  'T': 9,
  'Q': 10,
  'K': 11,
  'A': 12,
};

function getCardRank(card, partTwo = false) {
  return partTwo ? cardRanks2[card] : cardRanks[card];
}

const handTypes = {
  HIGH_CARD: 0,
  ONE_PAIR: 1,
  TWO_PAIR: 2,
  THREE_OF_A_KIND: 3,
  FULL_HOUSE: 4,
  FOUR_OF_A_KIND: 5,
  FIVE_OF_A_KIND: 6,
};


function getHandTypeKey(typeVal) {
  const entries = Object.entries(handTypes);
  return entries.find(e => e[1] === typeVal)[0];
}

function mapHands(hands) {
  return hands.map(handStr => {
    const [hand, bid] = handStr.split(" ");

    return {
      hand,
      typeRank: determineHandType(hand),
      get type() {
        return getHandTypeKey(this.typeRank);
      },
      bid: +bid,
      rank: -1,
    };
  });
}

function determineHandType(hand) {
  const matches = {};
  let highestMatchedCard = '';
  let highestMatchedNum = 0;

  hand.split("").forEach(card => {
    if (matches[card]) {
      matches[card] += 1;
    } else {
      matches[card] = 1;
    }

    if (card !== 'J' && matches[card] > highestMatchedNum) {
      highestMatchedNum = matches[card];
      highestMatchedCard = card;
    }
  });

  // Apply wilds to highest matched card
  if (matches['J'] && matches['J'] !== 5) {
    matches[highestMatchedCard] += matches['J'];
    matches['J'] = 0;
  }

  const matchCounts = Object.values(matches);

  switch (highestMatchedNum) {
    case 5:
      return handTypes.FIVE_OF_A_KIND;
    case 4:
      return handTypes.FOUR_OF_A_KIND;
    case 3:
      return matchCounts.includes(2)
        ? handTypes.FULL_HOUSE
        : handTypes.THREE_OF_A_KIND;
    case 2:
      return matchCounts.filter(mc => mc === 2).length === 1
        ? handTypes.ONE_PAIR
        : handTypes.TWO_PAIR;
    default:
      return handTypes.HIGH_CARD;
  }
}

function sortByCardStrength(cards, partTwo = false) {
  if (cards.length === 0) return [];
  if (cards.length === 1) return cards;

  let remaining = JSON.parse(JSON.stringify(cards));
  let current = remaining.shift();
  let sorted = [ current ];
  let cardIdx = 0;
  let idx = 0;
  let inserted = false;
  let skip = false;

  // While I have a card to grab and insert into sorted
  while (remaining.length) {
    current = remaining.shift();
    idx = 0;
    inserted = false;
     
    // While I haven't inserted my current card or gone to end of sorted
    while (!inserted && idx < sorted.length) {
      cardIdx = 0;
      skip = false;

      // While I haven't inserted, seen I can skip this hand, or gone through all cards in hand
      while (!inserted && !skip && cardIdx < HAND_LENGTH) {
        currentIsStronger = getCardRank(current.hand[cardIdx], partTwo)
          < getCardRank(sorted[idx].hand[cardIdx], partTwo);
        nextIsStronger = getCardRank(current.hand[cardIdx], partTwo)
          > getCardRank(sorted[idx].hand[cardIdx], partTwo);
        
        if (nextIsStronger) {
          sorted.splice(idx, 0, current);
          inserted = true;
        } else if (currentIsStronger) {
          skip = true;
        } else {
          cardIdx++;
        }
      }

      idx++;
    }

    if (!inserted) {
      sorted.push(current);
    }
  }

  return sorted.reverse(); // Return with weakest as first element
}

function main() {
  console.log("-- Day 7 --");

  const handsWithBids = getLines();
  const mappedHands = mapHands(handsWithBids);
  const sortedByTypeRank = mappedHands.sort((a, b) => a.typeRank - b.typeRank);
  const highestTypeRank = sortedByTypeRank[sortedByTypeRank.length - 1].typeRank;

  let currentTypeRank = 0;
  let currentTypeRankHands = [];
  let strengthRank = 1;

  const isPartTwo = true;

  while (currentTypeRank <= highestTypeRank) {
    currentTypeRankHands = sortedByTypeRank.filter(hand => hand.typeRank === currentTypeRank);

    if (currentTypeRankHands.length > 0) {
      const sortedByCardStrength = sortByCardStrength(currentTypeRankHands, isPartTwo);
      sortedByCardStrength.forEach(sortedHand => {
        currentTypeRankHands.find(h => h.hand === sortedHand.hand).rank = strengthRank;
        strengthRank++;
      });
    }

    currentTypeRank++;
  }
     
  const sortedByRank = sortedByTypeRank.sort((a, b) => a.rank - b.rank);
  const result = sortedByRank.reduce((acc, hand) => acc + (hand.bid * hand.rank), 0);
  console.log("P1 result:", result);
}

main();

