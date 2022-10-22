package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
	"unicode"
)

type Hangman struct {
	lentab          int    //longueur du tableau
	lenword         int    // variable qui stocke la longueur du mot choisi
	userletter      string // variable qui stocke la lettre entrée par l'utilisateur
	testsctword     string // variable qui stocke le mot choisi
	nbinwordpresent int    // variable qui stocke le nombre de fois que la lettre donné par l'utilisateur est présente dans le mot choisi
	slctletter      string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter1     string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter2     string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter3     string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter4     string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter5     string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter6     string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter7     string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter8     string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter9     string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter10    string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter11    string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter12    string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter13    string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter14    string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter15    string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	turn            int    // variable qui stocke le nombre de tour
	riddletab       []string
	indexuserletter int
	iswrong         bool
	attempt         int
	hasused         bool
}

func main() {
	var hangman1 Hangman
	hangman1.lentab = 0
	hangman1.lenword = 0
	hangman1.attempt = 10

	hangman("C:/Users/canel/Documents/Ynov/Tp/Hangman/word.txt", Hangman{lentab: 0, lenword: 0, attempt: 10})

}

func IsUpper(lettre string) bool {

	for _, letter := range lettre {
		if unicode.IsUpper(letter) {
			return true
		} else {
			fmt.Println("Votre lettre doit etre en majuscule")
			return false
		}
	}

	return false
}

func InputUser(hangman1 Hangman) {
	var Mo string

	fmt.Println("Choissisez une lettre : ")
	fmt.Scanln(&Mo)
	hangman1.userletter = Mo
	if len(Mo) == 1 {

		if IsUpper(Mo) == true {
			fmt.Println("Vous avez entré la lettre : ", Mo)
			lettercheck(hangman1)

		} else {
			InputUser(hangman1)
		}
	} else {
		fmt.Println("Vous devez entrer une seule lettre")
		InputUser(hangman1)
	}

}

func hangman(file string, hangman1 Hangman) {
	data, err := ioutil.ReadFile(file) // lire le fichier text.txt
	if err != nil {
		//	fmt.Println(fmt.Println("there is an error:", err))
		fmt.Println("_______________________________________________")
		fmt.Println("                                              ")
		fmt.Println("                                              ")
		fmt.Println("⚠️Erreur inattendue:  fermeture du programme⚠️")
		fmt.Println("                                              ")
		fmt.Println("_______________________________________________")

		fmt.Println("Cause:", err)
		fmt.Println()
		os.Exit(1)

	}

	//fmt.Println(string(data))
	//fmt.Println(len(data))
	var arr []string
	var tmp string

	for _, _byte := range data {

		if _byte == '\n' {
			arr = append(arr, tmp)
			tmp = ""
		} else {
			tmp += string(_byte)
		}
	}
	rand.Seed(time.Now().Unix())
	hangman1.lentab = len(arr)

	var randword = rand.Intn(hangman1.lentab)

	var selctword = arr[randword]
	hangman1.testsctword = selctword
	var lenselctword = len(arr[randword])

	for i := 0; i < len(selctword)-1; i++ {
		hangman1.lenword = i

	}

	hangman1.lenword = hangman1.lenword + lenselctword
	hangman1.lenword = hangman1.lenword / 2
	var randletter = rand.Intn(hangman1.lenword)
	hangman1.slctletter = string(selctword[randletter])

	//len(selctword) / 2
	//randletter = randletter - 1

	//

	var mysterytab []string
	for i := 0; i < len(selctword)-1; i++ {
		if i == randletter {
			mysterytab = append(mysterytab, string(selctword[randletter]))
		} else {
			mysterytab = append(mysterytab, "_")
		}
	}

	fmt.Println(mysterytab)
	hangman1.riddletab = mysterytab
	InputUser(hangman1)
}
func lettercheck(hangman1 Hangman) {

	var resultab []string
	for i := range hangman1.testsctword {
		if string(hangman1.testsctword[i]) == hangman1.userletter || string(hangman1.testsctword[i]) == hangman1.slctletter || string(hangman1.testsctword[i]) == hangman1.userletter1 || string(hangman1.testsctword[i]) == hangman1.userletter2 || string(hangman1.testsctword[i]) == hangman1.userletter3 || string(hangman1.testsctword[i]) == hangman1.userletter4 || string(hangman1.testsctword[i]) == hangman1.userletter5 || string(hangman1.testsctword[i]) == hangman1.userletter6 || string(hangman1.testsctword[i]) == hangman1.userletter7 || string(hangman1.testsctword[i]) == hangman1.userletter8 || string(hangman1.testsctword[i]) == hangman1.userletter9 || string(hangman1.testsctword[i]) == hangman1.userletter10 {
			resultab = append(resultab, string(hangman1.testsctword[i]))

			if string(hangman1.testsctword[i]) == hangman1.userletter {
				hangman1.nbinwordpresent = hangman1.nbinwordpresent + 1
			}
			hangman1.indexuserletter = i

		} else {
			resultab = append(resultab, "_")
		}
		if len(resultab) == len(hangman1.testsctword) {
			resultab = append(resultab[:i], resultab[i+1:]...)
		}
	}
	var iscorrect int
	for g := range resultab {
		if resultab[g] != "_" {

			iscorrect = iscorrect + 1

		}

	}
	if iscorrect >= len(resultab) {
		fmt.Println(resultab)
		win(hangman1)
	}

	if hangman1.nbinwordpresent >= 1 {
		if hangman1.userletter == hangman1.slctletter {
			if hangman1.nbinwordpresent == 1 {
				fmt.Println("❌La lettre est déja trouvée car donnée au début❌")
				hangman1.nbinwordpresent = 0
				hangman1.iswrong = true
				hangman1.attempt = hangman1.attempt - 1
				PositionHangman("hangman.txt", Hangman{attempt: hangman1.attempt, lentab: 0, lenword: 0})

				fmt.Println("Il vous reste", hangman1.attempt, "tentatives")
				InputUser(hangman1)
			}
			if hangman1.nbinwordpresent > 1 && hangman1.hasused == false {

				fmt.Println(resultab)
				hangman1.nbinwordpresent = 0
				hangman1.iswrong = false
				hangman1.hasused = true
				InputUser(hangman1)
			}
		} else {
			fmt.Println("✅La lettre est bien présente✅")
			hangman1.iswrong = false
			hangman1.turn = hangman1.turn + 1

			fmt.Print(resultab)
			switch hangman1.turn {
			case 1:
				{
					hangman1.userletter1 = hangman1.userletter
				}
			case 2:
				{
					hangman1.userletter2 = hangman1.userletter
				}
			case 3:
				{
					hangman1.userletter3 = hangman1.userletter
				}
			case 4:
				{
					hangman1.userletter4 = hangman1.userletter
				}
			case 5:
				{
					hangman1.userletter5 = hangman1.userletter
				}
			case 6:
				{
					hangman1.userletter6 = hangman1.userletter
				}
			case 7:
				{
					hangman1.userletter7 = hangman1.userletter
				}
			case 8:
				{
					hangman1.userletter8 = hangman1.userletter
				}
			case 9:
				{
					hangman1.userletter9 = hangman1.userletter
				}
			case 10:
				{
					hangman1.userletter10 = hangman1.userletter
				}
			case 11:
				{
					hangman1.userletter11 = hangman1.userletter
				}
			case 12:
				{
					hangman1.userletter12 = hangman1.userletter
				}
			case 13:
				{
					hangman1.userletter13 = hangman1.userletter
				}
			case 14:
				{
					hangman1.userletter14 = hangman1.userletter
				}
			case 15:
				{
					hangman1.userletter15 = hangman1.userletter
				}
			}

			hangman1.nbinwordpresent = 0
			InputUser(hangman1)
			//var isread bool

		}
	} else {
		fmt.Println("❌La lettre n'est pas dans le mot❌")
		hangman1.attempt = hangman1.attempt - 1
		fmt.Println("Il vous reste", hangman1.attempt, "tentatives")
		hangman1.iswrong = true
		PositionHangman("hangman.txt", Hangman{attempt: hangman1.attempt, lentab: 0, lenword: 0})

		InputUser(hangman1)
	}

	//copy(hangman1.riddletab, resultab)

	//fmt.Println(hangman1.iswrong)
	//fmt.Print(hangman1.riddletab, resultab)

	fmt.Print(resultab)

}
func PositionHangman(hangman string, hangman1 Hangman) { //fonction qui lie le fichier text des positions du hangman

	data, err := ioutil.ReadFile(hangman) // lire le fichier Position.txt
	if err != nil {                       // en cas d'erreur de liecture du fichier
		//fmt.Println(fmt.Println("there is an error:", err))
		fmt.Println("_______________________________________________")
		fmt.Println("                                              ")
		fmt.Println("                                              ")
		fmt.Println("⚠️Erreur inattendue:  fermeture du programme⚠️")
		fmt.Println("                                              ")
		fmt.Println("_______________________________________________")

		fmt.Println("Cause:", err)
		os.Exit(1)
	}

	var arr []string
	var tmp string

	for _, _byte := range data { // boucle qui permet d'ajouter ce qu'il y a dans le fichier texte dans le tableau

		if _byte == '.' {
			arr = append(arr, tmp)
			tmp = ""
		} else {
			tmp += string(_byte)
		}
	}

	fmt.Println(arr[hangman1.attempt])
	if hangman1.attempt == 0 {
		lose(hangman1)

	}
}

func win(hangman1 Hangman) {

	fmt.Println("VOUS AVEZ GAGNÉ  !!!!!!!!!!!!!")
	fmt.Println("le mot était", hangman1.testsctword)
	os.Exit(0)

}

func lose(hangman1 Hangman) {
	fmt.Println("VOUS AVEZ PERDU !!!!!!!!!!!!!")
	fmt.Println("le mot était", hangman1.testsctword)
	os.Exit(0)

}
