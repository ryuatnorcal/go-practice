package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {

	fmt.Println("Please enter your prices, Comfirm every price with enter")
	var prices []string
	for {
		var price string
		fmt.Print("Enter price: ")
		fmt.Scan(&price)
		prices = append(prices, price)
		if price == "0" {
			break
		}

	}
	return prices, nil
}
func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)

	return nil
}

func New() CMDManager {
	return CMDManager{}
}
