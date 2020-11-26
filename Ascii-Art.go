package main

import (
	//"github.com/01-edu/z01"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func getChar(style string) ([]rune, []rune) {
	filename := "ascii-table-template.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open " + filename + ": no such file or directory")
	}
	char := make([]byte, 283)
	file.Read(char)
	var runes []rune
	var ephemere []rune
	for i := 0; i < len(char); i++ {
		ephemere = append(ephemere, rune(char[i]))
	}
	for i := 0; i < len(ephemere); i++ {
		if ephemere[i] != 10 && ephemere[i] != 13 {
			runes = append(runes, rune(char[i]))
		}
	}
	file.Close()

	//initialisation du tableau suivant le style voulu

	var text []rune
	switch style {
	case "shadow":
		filename = "shadow.txt"
		file, err = os.Open(filename)
		if err != nil {
			fmt.Println("open " + filename + ": no such file or directory")
		}
		arr := make([]byte, 8317)
		file.Read(arr)
		for i := 0; i < len(arr); i++ {
			text = append(text, rune(arr[i]))
		}
		file.Close()
	case "thinkertoy":
		filename = "thinkertoy.txt"
		file, err = os.Open(filename)
		if err != nil {
			fmt.Println("open " + filename + ": no such file or directory")
		}
		arr := make([]byte, 5556)
		file.Read(arr)
		for i := 0; i < len(arr); i++ {
			text = append(text, rune(arr[i]))
		}
		file.Close()
	case "standart":
		filename = "standart.txt"
		file, err = os.Open(filename)
		if err != nil {
			fmt.Println("open " + filename + ": no such file or directory")
		}
		arr := make([]byte, 7474)
		file.Read(arr)
		for i := 0; i < len(arr); i++ {
			text = append(text, rune(arr[i]))
		}
		file.Close()
	default:
		filename = "standart.txt"
		file, err = os.Open(filename)
		if err != nil {
			fmt.Println("open " + filename + ": no such file or directory")
		}
		arr := make([]byte, 7474)
		file.Read(arr)
		for i := 0; i < len(arr); i++ {
			text = append(text, rune(arr[i]))
		}
		file.Close()

	}

	/*count:=0
	tab:= [][]rune{}
	for i := 0; i<855; i++{
		var row []rune
		for count != 7474{
			if text[count] != '\n' && text{
				row=append(row, text[count])
				fmt.Print(text[count])
			}else{
				break
			}
			count+=1
		}
		tab=append(tab, row)
	}*/
	return runes, text
}

func Split(s, sep string) ([]string, []int) {
	runes := []rune(s)
	seprunes := []rune(sep)
	position := []int{}
	position = append(position, 0)
	var tab []string
	lastpos := 0
	i := 0
	for i < len(runes) {
		if runes[i] == seprunes[0] {
			if string(runes[i:i+len(seprunes)]) == sep {
				position = append(position, i+2)
				tab = append(tab, string(runes[lastpos:i]))
				lastpos = i + len(seprunes)
			}
		}
		i++
	}
	if lastpos != i {
		tab = append(tab, string(runes[lastpos:i]))
	}
	position = append(position, len(os.Args[1]))
	return tab, position
}

func Colors(tabColor string) string {
	Couleur := "\033[0m"
	color := tabColor
	if color != "rien" {
		couleurs := color[8:]
		Couleur = "\033[0m"
		switch couleurs {
		case "Red":
			Couleur = "\033[31m"
		case "Green":
			Couleur = "\033[32m"
		case "Yellow":
			Couleur = "\033[33m"
		case "Blue":
			Couleur = "\033[34m"
		case "Purple":
			Couleur = "\033[35m"
		case "Cyan":
			Couleur = "\033[36m"
		case "White":
			Couleur = "\033[37m"
		default:
			Couleur = "\033[0m"
		}
	} else {
		Couleur = "\033[0m"
	}
	return Couleur
}

func Analyse() ([]string, []string, []string, []string, []string) {
	argument := (os.Args[1:])
	var texte []string
	var style []string
	var color []string
	var output []string
	var justify []string
	for i := 0; i < len(argument); i++ {
		argumentTest := argument[i]
		if len(argumentTest) > 9 && argumentTest[:8] == "--color=" {
			for len(texte) > len(color)+1 {
				color = append(color, "rien")
			}
			color = append(color, argument[i])
		} else if argumentTest == "shadow" || argumentTest == "thinkertoy" || argumentTest == "standart" {
			for len(texte) > len(style)+1 {
				style = append(style, "rien")
			}
			style = append(style, argument[i])
		} else if len(argumentTest) > 10 && argumentTest[:9] == "--output=" {
			for len(texte) > len(output)+1 {
				output = append(output, "rien")
			}
			output = append(output, argument[i])
		} else if len(argumentTest) > 9 && argumentTest[:8] == "--align=" {
			for len(texte) > len(justify)+1 {
				justify = append(justify, "rien")
			}
			justify = append(justify, argument[i])
		} else {
			texte = append(texte, argument[i])
		}
	}
	for len(style) != len(texte) {
		style = append(style, " ")
	}
	for len(color) != len(texte) {
		color = append(color, "rien")
	}
	for len(output) != len(texte) {
		output = append(output, "rien")
	}
	for len(justify) != len(texte) {
		justify = append(justify, "rien")
	}
	return texte, style, color, output, justify
}

func size() []rune {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	String := string(out)
	register := false
	sorti := []rune{}
	for i := 0; i < len(String); i++ {
		if register && String[i] != 10 {
			sorti = append(sorti, rune(String[i]))
		}
		if String[i] == 32 {
			register = true
		}
	}
	return sorti
}

func main() {
	T := size()
	terminalSize := 0
	for i := 0; i < len(T); i++ {
		op := int(T[i]) - 48
		terminalSize *= 10
		terminalSize += op
	}
	tabTexte, tabStyle, tabColor, tabOut, tabJustify := Analyse()

	for mot := 0; mot < len(tabTexte); mot++ {
		style := ""
		style = tabStyle[mot]
		args := tabTexte[mot]
		var Str string
		for i := 0; i < len(args); i++ {
			Str += string(args[i])
		}
		tab, position := Split(Str, "\\n")
		runes, text := getChar(style)

		//Replacement des mots
		/* for i:=0; i<len(position)-2; i++{
			sep:=position[i+1]
			for change:=sep-1; change<len(tab); change++{
				tab[change-1]=tab[change]
			}
		} */

		part := len(position) - 2
		for word := 0; word <= part; word++ {
			print := [8][]string{}
			//taille:=len(tab[word])

			for _, i := range tab[word] {
				pos := 0
				countn := 0
				if i == '~' {
					for ligne := 0; ligne < 8; ligne++ {

						for countn != 848+ligne-2 {
							pos++
							if text[pos] == '\n' {
								countn++
							}
						}
						point := []string{" "}
						test := []string{""}
						count := 0
						for test[count] != "\n" {
							count++
							result := string(text[count+pos])
							test = append(test, result)
						}
						for copie := 0; copie < len(test)-2; copie++ {
							point = append(point, test[copie])
						}
						for l := 1; l < len(point); l++ {
							print[ligne] = append(print[ligne], point[l])
						}
					}
				} else {
					for j := 0; j < len(runes); j++ {
						if i == (runes[j]) {
							//9*j -7
							for ligne := 0; ligne < 8; ligne++ {
								pos := 0
								countn := 0
								if style == "thinkertoy" {
									for countn != j*9+ligne+2 {
										pos++
										if text[pos] == '\n' {
											countn++
										}
									}
								} else {
									//cherche la position du texte avec l'aide de l'ascii table
									for countn != j*9+ligne+1 {
										pos++
										if text[pos] == '\n' {
											countn++
										}
									}
								}
								point := []string{" "}
								test := []string{""}
								count := 0
								for test[count] != "\n" {
									count++
									result := string(text[count+pos])
									test = append(test, result)
								}
								for copie := 0; copie < len(test)-2; copie++ {
									point = append(point, test[copie])
								}
								for l := 1; l < len(point); l++ {
									//fmt.Print(point[l])
									print[ligne] = append(print[ligne], point[l])
								}
							}
						}
					}
				}
			}

			eph := tabJustify[mot]
			var align string
			if eph != "rien" {
				align = eph[8:]
			}
			//affichage
			if align == "right" {

				for ligne := 0; ligne < len(print); ligne++ {
					nbEspace := terminalSize - len(print[ligne])
					for espace := 0; espace < nbEspace; espace++ {
						fmt.Print(" ")
					}
					for i := 0; i < len(print[ligne]); i++ {
						couleur := Colors(tabColor[mot])
						fmt.Print(string(couleur))
						if print[ligne][i] != "\n" && i != len(print[ligne]) {
							fmt.Print(string(couleur), print[ligne][i])
						}
					}
					fmt.Println()
				}
			} else if align == "center" {
				for ligne := 0; ligne < len(print); ligne++ {
					nbEspace := (terminalSize - len(print[ligne])) / 2
					for espace := 0; espace < nbEspace; espace++ {
						fmt.Print(" ")
					}
					for i := 0; i < len(print[ligne]); i++ {
						couleur := Colors(tabColor[mot])
						fmt.Print(string(couleur))
						if print[ligne][i] != "\n" && i != len(print[ligne]) {
							fmt.Print(string(couleur), print[ligne][i])
						}
					}
					fmt.Println()
				}
			} else {
				for ligne := 0; ligne < len(print); ligne++ {
					for i := 0; i < len(print[ligne]); i++ {
						//countSpace:=0
						couleur := Colors(tabColor[mot])
						fmt.Print(string(couleur))
						if print[ligne][i] != "\n" && i != len(print[ligne]) {
							fmt.Print(string(couleur), print[ligne][i])
						}
					}
					fmt.Println()
				}
			}

			output := tabOut[mot]
			if len(output) > 9 && output[0:9] == "--output=" {
				output = output[9:len(output)]
				f, err := os.Create(output)
				if err != nil {
					fmt.Print("Error occured")
				}

				defer f.Close()

				for ligne := 0; ligne < len(print); ligne++ {
					for i := 0; i < len(print[ligne]); i++ {
						if print[ligne][i] != "\n" && i != len(print[ligne]) {
							_, err = f.WriteString(print[ligne][i])
							if err != nil {
								fmt.Print("Error occured")
							}
						}
					}
					_, err = f.WriteString("\n")
				}
			}

		}
	}
}
