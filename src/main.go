package main

// so da para atribuir valor dentro de um code block:
// var x int
// x = 10 ESTA ERRADA
// func main() {
// 	x = 10 // SO FUNCIONA SE TIVER DENTRO DO CODE BLOCK
// 	fmt.Println(x)
//}
// o ato de definir, criar, estruturar tipos compostos chamase-se composição.

// tipo em GO são estáticos
// := declaração curta de variável so pode ser usada dentro de funções/code blocks
// var declaração normal de variável pode ser usada em qualquer lugar do código
import (
	//"bufio"
	"fmt"
	//"os"
	//"os/exec"
	//"strings"
)

var y int = 10

func main() {
	z := 20
	qualquercoisa(z)

	// marmota sempre dentro dos metodos.
	// var e fora. packet level scope
}
func qualquercoisa(x int) {
	fmt.Println(y)
	fmt.Println(x)
}

// fmt.Print("Deseja abrir o Bloco de Notas? (s/n): ")
// reader := bufio.NewReader(os.Stdin)
// resposta, _ := reader.ReadString('\n')
// resposta = strings.TrimSpace(resposta)

// if resposta == "s" || resposta == "S" {
//fmt.Println("Abrindo o Bloco de Notas...")
//cmd := exec.Command("notepad.exe")
//cmd.Start()
//} else {
//fmt.Println("Operação cancelada.")
//}
//x := 10
//y := "apenas"
//fmt.Printf("x: %v, y: %v", x, x)
//fmt.Printf("y: %v, x: %v", y, y)
// %v : para valor
// %T : para tipo da variável
// %d : para números decimais
// %b : para números binários
// %#u : para unicode
// %#x : para números hexadecimal
