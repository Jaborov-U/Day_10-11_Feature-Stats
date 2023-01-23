package stats

import (
	"fmt"

	"github.com/Jaborov-U/Day_10-11_Feature-Stats/pkg/bank/types"
)

func ExampleAVG() {

	cards := []types.Payment{
		{
			Amount: 10,
		},
		{
			Amount: 15,
		},
		{
			Amount: 20,
		},
	}
	avgPays := AVG(cards)

	fmt.Println(avgPays)

	//Output:15

}
