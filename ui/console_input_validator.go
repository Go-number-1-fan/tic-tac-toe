package ui

type ConsoleInputValidator struct{}

func (validator ConsoleInputValidator) IsValid(selection int, validSelections []int) bool {
	for _, validSelection := range validSelections {
		if selection == validSelection {
			return true
		}
	}
	return false
}
