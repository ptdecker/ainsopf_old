package profile

import (
    "fmt"
)

var cardNames = [...]string{
    "The Fool",
    "The Magician",
    "The High Priestess",
    "The Empress",
    "The Emperor",
    "The Hierophant",
    "The Lovers",
    "The Chariot",
    "Justice",
    "The Hermit",
    "Wheel of Fortune",
    "Strength",
    "The Hanged Man",
    "Death",
    "Temperance",
    "The Devil",
    "The Tower",
    "The Star",
    "The Moon",
    "The Sun",
    "The Last Judgement",
    "The World"}

func sumDigits(num int) int {
    sumOfDigits := 0
    for ; num > 0; num /= 10 {
        sumOfDigits += num % 10
    }
    return sumOfDigits
}

func buildProfile(birthYear, fromYear, toYear int) {
    for year := fromYear; year < (toYear + 1); year++ {
        value := sumDigits(year) + sumDigits(year - birthYear)
        if value > 22 {
            value = sumDigits(value)
        }
        if value == 22 {
            value = 0
        }
        fmt.Printf("%4v %2v %s\n", year, value, cardNames[value])
    }
}

func Gen() {
    buildProfile(1984, 1984, 2084)
}

