package main

// := declaração curta de variável so pode ser usada dentro de funções/code blocks
// var declaração normal de variável pode ser usada em qualquer lugar do código
import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Print("Deseja abrir o Bloco de Notas? (s/n): ")
	reader := bufio.NewReader(os.Stdin)
	resposta, _ := reader.ReadString('\n')
	resposta = strings.TrimSpace(resposta)

	if resposta == "s" || resposta == "S" {
		fmt.Println("Abrindo o Bloco de Notas...")
		cmd := exec.Command("notepad.exe")
		cmd.Start()
	} else {
		fmt.Println("Operação cancelada.")
	}
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

}
