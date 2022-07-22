/*
Copyright © 2022 samjtro

*/
package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var lowercase = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var uppercase = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var numbers = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
var special = []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '-', '+', '=', '{', '[', '}', ']', '|', ';', ':', ',', '.', '<', '>', '/', '?'}

var lLength = len(lowercase)
var uLength = len(uppercase)
var nLength = len(numbers)
var sLength = len(special)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "New Password",
	Long: `pass new <len> <spec>
len = Length of the password requested. Must be > 8.
spec = Whether or not to include special characters in the password requested. Must be true or false. 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		var password string
		var rng int
		len, err := strconv.Atoi(args[0])

		if err != nil {
			log.Fatalf(err.Error())
		} else if len <= 8 {
			log.Fatalf("Insecure password! Please choose a length longer than 8.")
		}

		for i := 0; i < len; i++ {
			rng = rand.Intn(3)
			switch args[1] {
			case "true":
				switch rng {
				case 0:
					password += string(lowercase[rand.Intn(lLength)])
				case 1:
					password += string(uppercase[rand.Intn(uLength)])
				case 2:
					password += string(numbers[rand.Intn(nLength)])
				case 3:
					password += string(special[rand.Intn(sLength)])
				}
			case "false":
				switch rng {
				case 0:
					password += string(lowercase[rand.Intn(lLength)])
				case 1:
					password += string(uppercase[rand.Intn(uLength)])
				case 2:
					password += string(numbers[rand.Intn(nLength)])
				}
			default:
				log.Fatalf("Must choose true or false for spec.")
			}
		}

		fmt.Printf("Your Password is Below:\n%s\n", password)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
