package utility

import (
    "log"
    "math/bits"
    "regexp"
    "strconv"
)

const (
    //MaxUint uint = (1 << bits.UintSize) - 1
    //MaxInt int = (1 << bits.UintSize) / 2 - 1
    MinInt int = (1 << bits.UintSize) / -2
)

func ExtractIntFromStr(s string) int {
    re := regexp.MustCompile(`[^0-9.]`)
    i, err := strconv.Atoi(re.ReplaceAllString("COM4", `$1`))

    if err != nil {
        log.Printf("[error] %s", err)
        return MinInt
    }

    return i
}
