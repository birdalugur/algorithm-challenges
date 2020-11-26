package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the countApplesAndOranges function below.
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
        stApple := 0
        stOrange := 0
        for i := 0; i < len(apples); i++ {
            apples[i] += a
            if apples[i] <= t && apples[i] >= s {
                stApple++
            }
        }
        for i := 0; i < len(oranges); i++ {
            oranges[i] += b
            if oranges[i] <= t && oranges[i] >= s {
                stOrange++
            }
        }
        
        fmt.Println(stApple)
        fmt.Println(stOrange)

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    st := strings.Split(readLine(reader), " ")

    sTemp, err := strconv.ParseInt(st[0], 10, 64)
    checkError(err)
    s := int32(sTemp)

    tTemp, err := strconv.ParseInt(st[1], 10, 64)
    checkError(err)
    t := int32(tTemp)

    ab := strings.Split(readLine(reader), " ")

    aTemp, err := strconv.ParseInt(ab[0], 10, 64)
    checkError(err)
    a := int32(aTemp)

    bTemp, err := strconv.ParseInt(ab[1], 10, 64)
    checkError(err)
    b := int32(bTemp)

    mn := strings.Split(readLine(reader), " ")

    mTemp, err := strconv.ParseInt(mn[0], 10, 64)
    checkError(err)
    m := int32(mTemp)

    nTemp, err := strconv.ParseInt(mn[1], 10, 64)
    checkError(err)
    n := int32(nTemp)

    applesTemp := strings.Split(readLine(reader), " ")

    var apples []int32

    for i := 0; i < int(m); i++ {
        applesItemTemp, err := strconv.ParseInt(applesTemp[i], 10, 64)
        checkError(err)
        applesItem := int32(applesItemTemp)
        apples = append(apples, applesItem)
    }

    orangesTemp := strings.Split(readLine(reader), " ")

    var oranges []int32

    for i := 0; i < int(n); i++ {
        orangesItemTemp, err := strconv.ParseInt(orangesTemp[i], 10, 64)
        checkError(err)
        orangesItem := int32(orangesItemTemp)
        oranges = append(oranges, orangesItem)
    }

    countApplesAndOranges(s, t, a, b, apples, oranges)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
