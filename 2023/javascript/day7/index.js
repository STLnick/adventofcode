const fs = require("fs");

function getLines() {
    const filename = process.argv.includes("t") || process.argv.includes("-t")
        ? "input-test.txt"
        : "input.txt";
    return fs.readFileSync(__dirname + "/" + filename).toString().split("\n").filter(Boolean);
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

        // TODO sort cards by strength

        // TODO find high card
       
        // TODO find type
        
        return {
            hand,
            highCard,
            type,
            bid: +bid,
            rank: -1, // will rank later
        };
    });
}

function determineHandType(hand) {
    let type = '';
    const firstCard = hand[0];
    let matches = 1;

    for (let i = 1; i < hand.length; i++) {
        if (hand[i] === firstCard) {
            matches++;
        }
    }

    switch (matches) {
        case 5:
            return handTypes.FIVE_OF_A_KIND;
        case 4:
            return handTypes.FOUR_OF_A_KIND;
        case 3:
            // 3 of a kind
            // OR full house
        case 2:
            // two pair or one pair
        case 1:
    }

    // five of a kind
    // four of a kind
    // full house
    // three of a kind
    // two pair
    // one pair
    // high card

    return type;
}

function main() {
    console.log("-- Day 7 --");

    const lines = getLines();
    console.log(lines);
}

main();

