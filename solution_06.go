package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
    length := len(arr)
    l := float32(length)
    
    var pozitive, negative, zero float32
    var pRatio, nRatio, zRatio float32
    
    pozitive,negative,zero = 0,0,0
    
    for _, item := range arr{
        if item>0{
            pozitive++
        }else if item<0{
            negative++
        }else{
            zero++
        }
    }
    pRatio, nRatio, zRatio = pozitive/l, negative/l, zero/l
    
    
    fmt.Printf("%.6f\n", pRatio) 
    fmt.Printf("%.6f\n", nRatio) 
    fmt.Printf("%.6f\n", zRatio) 
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
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
