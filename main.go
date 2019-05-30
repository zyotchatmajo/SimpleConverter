package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func cToJava(str string) string {
	str = TabToSpace(str)
	a := strings.Split(str, "(") // separa la string
	if strings.Contains(str, "int main()"){
		return "public static void main(String ar[]){"
	}
	if strings.Contains(str, "printf") {
		//si tiene variable separa la string por pedasos y lo junta todo
		if strings.Contains(str, "%d") {
			ArrayS := strings.Split(str, "%d")
			ArrayS2 := strings.Split(str, ",")
			ArrayS3 := strings.Split(ArrayS2[1], ")")
			LastPart := strings.Split(ArrayS[1], ",")
			if LastPart[0] == string(34) {
				LastPart[0] = ""
				return "System.out.print" + strings.ReplaceAll(ArrayS[0], "printf", "") + "" + string(34) + "+" + ArrayS3[0] + ");"
			}else{
				return "System.out.print" + strings.ReplaceAll(ArrayS[0], "printf", "") + "" + string(34) + "+" + ArrayS3[0] + "+" + string(34) + "" + LastPart[0] + ");"
			}
			
		}
		// simplemente regresa la impresion
		return "System.out.print(" + a[1]
		//fmt.Println("Printf Function")
	}
	if strings.Contains(str, "scanf") {
		//separa la variable y dependiendo de si es d o f asigna int o float
		var variable strings.Builder
		var x int
		min := strings.Index(str, "&")
		max := strings.Index(str, ")")
		for x = min + 1; x < max; x++ {
			variable.WriteByte(str[x])
		}
		if strings.Contains(str, "%d"){
			return variable.String() + " = teclado.nextInt();"
		}else if strings.Contains(str,"%f"){
			return variable.String() + " = teclado.nextFloat();"
		}

	}
	return str
}

func JavaToC(str string) string {
	str = TabToSpace(str)
	if strings.Contains(str,"public static void main"){
		return "int main(){"
	}
	if strings.Contains(str,"Scanner"){
		return ""
	}
	//Identifica si la instruccion es una impresion en consola
	if strings.Contains(str, "System.out.println") {
		if strings.Contains(str, "+") {
			//Separa el string para que sea mas facil manipularlo
			ArrayS := strings.Split(str, "(")
			//Separa el string en texto y variable
			ArrayS2 := strings.Split(ArrayS[1], "+")
			//Remplaza caracteres que interfieren en la sintaxis
			ArrayS2[1] = strings.ReplaceAll(ArrayS2[1], ")", "")
			ArrayS2[1] = strings.ReplaceAll(ArrayS2[1], ";", "")
			ArrayS2[1] = strings.ReplaceAll(ArrayS2[1], "\"", "")
			ArrayS2[0] = strings.ReplaceAll(ArrayS2[0], "\"", "")
			//regresa un string ya con la sintaxis correcta
			return "printf(\"" + ArrayS2[0] + "%d\"," + ArrayS2[1] + ");"
		} else {
			ArrayS := strings.Split(str, "\"")
			return "printf(\"" + ArrayS[1] + "\");"
			//separa el string para que solo se pueda utilizar lo que esta dentro de las comillas
		}
	}
	//Scanf diferencia entre INT y FLOAT
	if strings.Contains(str, "teclado.nextInt()"){
		ArrayS := strings.Split(str, "=")
		return "scanf("+ string(34) + "%d"+ string(34) + ",&"+ArrayS[0]+");"
	}
	if strings.Contains(str, "teclado.nextFloat()"){
		ArrayS := strings.Split(str, "=")
		return "scanf("+ string(34) + "%f"+ string(34) + ",&"+ArrayS[0]+");"
	}
	return str
}
func main() {
	//Menu 
	fmt.Println("1. Java a C 2. C a Java")
	var input int 
	scannerb := false
	fmt.Scanln(&input)
	//Abrir archivo y crear
	file, err := os.Open("file.txt")
	f, erro := os.Create("test.txt")
	if erro != nil {
        fmt.Println(erro)
        return
    }
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(JavaToC(scanner.Text())) debug
		if input == 1 { 		//lectura para java
			l, erro := f.WriteString(JavaToC(scanner.Text()) + "\n")
			fmt.Println(l, "bytes written successfully")
			if erro != nil {
				fmt.Println(erro)
				f.Close()
				return
			} 
		}else if input == 2 { 		//lectura para c
			l, erro := f.WriteString(cToJava(scanner.Text()) + "\n")
			fmt.Println(l, "bytes written successfully")
			if erro != nil {
				fmt.Println(erro)
				f.Close()
				return
			}
			if scannerb == false { 		//introducir scanner
				l, erro := f.WriteString("Scanner teclado = new Scanner(System.in);" + "\n")
				fmt.Println(l, "bytes written successfully")
				if erro != nil {
					fmt.Println(erro)
					f.Close()
					return
				}
				scannerb = true
			}
		}
	}
    erro = f.Close()
    if erro != nil {
        fmt.Println(erro)
        return
    }
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
//Remueve tabs pero alparecer no funciona :(
func TabToSpace(input string) string {
	var result []string

	for _, i := range input {
			switch {
			// all these considered as space, including tab \t
			// '\t', '\n', '\v', '\f', '\r',' ', 0x85, 0xA0
			case unicode.IsSpace(i):
					result = append(result, " ") // replace tab with space
			case !unicode.IsSpace(i):
					result = append(result, string(i))
			}
	}
	return strings.Join(result, "")
}