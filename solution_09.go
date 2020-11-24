package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(date string) string {
    /*
     * Write your code here.
     */
    var hh, mm, ss int

    hh, _ = strconv.Atoi(date[0:2])
    mm, _ = strconv.Atoi(date[3:5])
    ss, _ = strconv.Atoi(date[6:8])
    period := date[len(date)-2:len(date)]

    if period == "PM" && hh != 12 {
        hh += 12
    }else if period == "AM" && hh == 12{
        hh -= 12
    }
    
    res := fmt.Sprintf("%02v:%02v:%02v",hh,mm,ss)

    return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer outputFile.Close()

    writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
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
