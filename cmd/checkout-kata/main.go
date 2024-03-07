package main

import (
	"bufio"
	"fmt"
	"os"

	"checkout-kata/internal"
)

const exitCommand string = "quit"
const filePath string = "./products_prices.yaml"

func main() {
	priceParser := internal.Parser{}
	pricing := internal.NewPricing(&priceParser, filePath)
	err := pricing.InitPrices()
	if err != nil {
		fmt.Println("Error initializing prices:", err)
		return
	}
	checkout := internal.NewCheckout(pricing)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Helper commands:\n - quit: exit the program\n ----------")

	for {
		fmt.Print("Enter a product SKU: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		line = line[:len(line)-1] // remove newline

		if len(line) == len(exitCommand) && line == exitCommand {
			fmt.Println("Exiting...")
			return
		}

		checkout.Scan(line)
		fmt.Println(checkout.String())
		fmt.Printf("Total price: %d\n ---------- \n", checkout.GetTotalPrice())
	}
}
