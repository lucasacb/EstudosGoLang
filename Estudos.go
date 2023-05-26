// declaração de pacote principal
package main

//importar função fmt
import "fmt"

//declaração da variavel CONST do ponto de ebulição da água em F
const ebulicaoK int  = 373

//função principal
func main () {

	tempK := ebulicaoK //Variavél do valor da temperatura em graus K - Declarando usando o operador curto ( := )
	var tempC  = (tempK - 273) //Conversão de K para C - Declarando usando o var

	//fmt.Println("A temperatura de ebulição da água em ºK =", tempK)
	//fmt.Println("A temperatura de ebulição da água em ºC =", tempC)
	fmt.Printf(" A temperatura de ebulição da água em ºK = %v e A temperatura de ebulição da água em ºC = %v", tempK, tempC)

}