package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    var instr string

    if len(os.Args)>1{
        instr = strings.ToUpper(os.Args[1])
        fmt.Println(instr)
        outstr := hexToBin(instr)
        fmt.Println(outstr)
        num, err := strconv.ParseInt(outstr, 2, 64)
        if err != nil {
           panic(err)
        }
        fmt.Println("decimal num: ", num)
    }
}

func hexToBin(inHex string) (string) {
    var outBin string
    outBin = ""
    for i := 0; i < len(inHex); i++ {
        char := string(inHex[i])
        println(char)
        outstr := hexCharToBinChar(char)
        outBin += outstr
    }
    return outBin
}

func hexCharToBinChar(charIn string) (string) {
    var charOut string
    switch charIn {
       case "0": charOut = "0000"
       case "1": charOut = "0001"
       case "2": charOut = "0010"
       case "3": charOut = "0011"
       case "4": charOut = "0100"
       case "5": charOut = "0101"
       case "6": charOut = "0110"
       case "7": charOut = "0111"
       case "8": charOut = "1000"
       case "9": charOut = "1001"
       case "A": charOut = "1010"
       case "B": charOut = "1011"
       case "C": charOut = "1100"
       case "D": charOut = "1101"
       case "E": charOut = "1110"
       case "F": charOut = "1111"
    }
    return charOut
}
