package main

import (
    "os"
    "bufio"
	"fmt"
    "slices"
    "strings"
    "strconv"
)

func main() {
    part1()
    // part2()
}

var cards = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
const (
    HighCard = iota
    OnePair
    TwoPair
    ThreeKind
    FullHouse
    FourKind
    FiveKind
)

func getHandType(h string) int {
    // h is five card hand, like "AK134"
    occ := make(map[rune]int)
    for _, c := range h {
        occ[c]++
    }
    // Determine type
    v := []int{}
    for _, o := range occ {
        v = append(v, o)
    }
    switch len(occ) {
    case 5:
        return HighCard
    case 4:
        return OnePair
    case 3:
        // Two Pairs or Three of a kind
        if slices.Contains(v, 3) {
            return ThreeKind
        }
        return TwoPair
    case 2:
        // Full House or Four of kind
        if slices.Contains(v, 4) {
            return FourKind
        }
        return FullHouse
    case 1:
        return FiveKind
    }
    return 0
}

type Hand struct {
    hand string
    bid int
}

func part1() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)
    rankedDesc := []Hand{}
    // Foreach hand (line in input)
    for scanner.Scan() {
        bid, _ := strconv.Atoi(strings.Fields(scanner.Text())[1])
        h := Hand{
            strings.Fields(scanner.Text())[0],
            bid,
        }
        if len(rankedDesc) == 0 {
            rankedDesc = append(rankedDesc, h)
            continue
        }
        for i, rankedHand := range rankedDesc {
            if getHandType(h.hand) > getHandType(rankedHand.hand) {
                rankedDesc = slices.Insert(rankedDesc, i, h)
                break
            }

            if getHandType(h.hand) < getHandType(rankedHand.hand) && i < len(rankedDesc) - 1 {
                continue
            }

            if getHandType(h.hand) == getHandType(rankedHand.hand) {
                added := false
                for j := 0; j < len(h.hand); j++ {
                    if slices.Index(cards, rune(h.hand[j])) == slices.Index(cards, rune(rankedHand.hand[j])) {
                        continue
                    }
                    if slices.Index(cards, rune(h.hand[j])) > slices.Index(cards, rune(rankedHand.hand[j])) {
                        rankedDesc = slices.Insert(rankedDesc, i, h)
                        added = true
                        break
                    }
                }
                if added {
                    break
                }
            }
            rankedDesc = append(rankedDesc, h)
            // rankedDesc = slices.Insert(rankedDesc, i+1, h)
            break
        }
    }

    totalWin := 0
    for i := len(rankedDesc); i > 0; i-- {
        totalWin += rankedDesc[i - 1].bid * i
    }
    fmt.Println("----- The Answer My Friend is : ", totalWin, " -----")
}

func part2() {
}

