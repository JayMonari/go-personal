package gross

// Units store the Gross Store unit measurement
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill create a new bill
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem add item to customer bill
func AddItem(bill, units map[string]int, item, unit string) bool {
	if val, ok := units[unit]; ok {
		bill[item] = val
		return true
	}
	return false
}

// RemoveItem remove item from customer bill
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	u, uok := units[unit]
	i, iok := bill[item]
	if !uok || !iok {
		return false
	} else if i-u < 0 {
		return false
	}

	bill[item] = i - u
	if bill[item] == 0 {
		delete(bill, item)
	}

	return true
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
	if v, ok := bill[item]; ok {
		return v, true
	}
	return 0, false
}
