package stats

import "github.com/Jaborov-U/Day_10-11_Feature-Stats/pkg/bank/types"

// AVG рассчитывает среднюю сумму платежа.
func AVG(payments []types.Payment) types.Money {

	summPays := 0

	for _, payment := range payments {
		summPays += int(payment.Amount)
	}

	return types.Money(summPays / len(payments))
}

// TotalInCategory находит сумму покупок в определенной категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money{
	
	summPays := 0

	for _, payment := range payments {
		summPays += int(payment.Amount)
		category += payment.Category
	}

	return types.Money(summPays / len(payments))

}