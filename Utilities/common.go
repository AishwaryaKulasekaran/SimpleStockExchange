package utilities

import (
	model "SimpleStockExchange/Models"
	"fmt"
)

// removeElement deletes the element at index `i` and preserves the order of the slice `s`.The original slice is not modified here
func RemoveElement(s []model.Order, i int) ([]model.Order, error) {

	// perform bounds checking first to prevent a panic!
	if i >= len(s) || i < 0 {
		return nil, fmt.Errorf("Index is out of range. Index is %d with slice length %d", i, len(s))
	}
	newSlice := make([]model.Order, 0)
	newSlice = append(newSlice, s[:i]...)

	return append(newSlice, s[i+1:]...), nil
}
