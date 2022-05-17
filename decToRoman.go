package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func TradMillions(number int) string {
	var tradMillion string
	for i:= 0; i < number; i++ {
		tradMillion += "M"
	}
	tradMillion += "|"
	return tradMillion
}
func TradHundreds(number int) string {
	var tradHundreds [9]string = [9]string{"C","CC", "CCC", "CD", "D", "DC", "DCC","DCCC","CM"}
	return tradHundreds[number - 1]
}
func TradTens(number int) string {
	var tradTens [9]string = [9]string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX","XC"}
	return tradTens[number - 1]
}
func TradDec(number int)string {
	var tradDec [9]string = [9]string{"I", "II", "III", "IV", "V","VI","VII","VIII", "IX"}
	return tradDec[number - 1 ]
}
func ConvertDecToRom(number int) string{
	var Rom string
	if number / 1000000 != 0{
		numberMillion := number / 1000000
		Rom += TradMillions(numberMillion)
		numberThousands := number % 1000000
		if numberThousands != 0 {
			Rom += ConvertDecToRom(numberThousands)
		}
	}else if number / 1000 != 0 {
		numberThousands := number / 1000
		Rom += ConvertDecToRom(numberThousands) + "|"
		numberHundreds := number % 1000
		if numberHundreds != 0 {
			Rom += ConvertDecToRom(numberHundreds)
		}
	}else if number / 100 != 0 {
		numberHundreds := number / 100
		Rom += TradHundreds(numberHundreds)
		numberTens := number % 100
		if numberTens != 0 {
			Rom += ConvertDecToRom(numberTens)
		}
	} else if number / 10 != 0{
		numberTens := number / 10
		Rom += TradTens(numberTens)
		numberDec := number % 10
		if numberDec != 0 {
			Rom += ConvertDecToRom(numberDec)
		}
	}else {
		Rom += TradDec(number)
	}
	return Rom
}
func main(){
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese el numero: ")
	reader.Scan()
	input := reader.Text()
	inputINT, err := strconv.Atoi(input)
	if err != nil || inputINT < 1 {
		fmt.Println("Error")
	}else {
		rom := ConvertDecToRom(inputINT)
		fmt.Println("El numero equivalete de " + input + " es " + rom)
	}


}
