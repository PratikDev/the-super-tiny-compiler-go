package tokenizer

import "regexp"

// Token is a struct that represents a token
type Token struct {
	Type  string
	Value string
}

// tokenizer is a function that takes a string input and returns a slice of tokens
// It uses a regular expression to match different types of tokens
// such as parentheses, numbers, strings, and names
func Tokenizer(input string) []Token {
	tokens := []Token{}

	for i := 0; i < len(input); {
		current := string(input[i])

		// if string is a parenthesis start or end
		if current == "(" || current == ")" {
			tokens = append(tokens, Token{
				Type:  "paren",
				Value: current,
			})

			i++
			continue
		}

		rgBlank := regexp.MustCompile(`\s`) // regex for blank spaces
		// if we find some blank space
		// they can fuck themselves. we don't care about them
		if rgBlank.MatchString(current) {
			i++
			continue
		}

		rgDigit := regexp.MustCompile(`[0-9]`) // regex for digists
		/** if we find a digit,
		* we can't only push the current value
		* to the `tokens` array. we need to cover the
		* whole value of the number like "1234" instead
		* of only "1"
		**/
		if rgDigit.MatchString(current) {
			// getting the value of `current` in `value` store
			value := current

			// increasing `i` here cuz the last element
			// have been checked by the above `if`
			i++

			// to cover the whole number, we do
			// a loop until we find a blank space
			// (getting element manually so that it gets the new
			// `i` in each loop)
			for !rgBlank.MatchString(string(input[i])) && string(input[i]) != ")" {
				// for easy access
				current = string(input[i])

				// if we have a non-number along with numbers
				if !rgDigit.MatchString(current) {
					panic("invalid number syntax")
				}

				value += current

				i++
				continue
			}

			// all numerical values have been processed
			tokens = append(tokens, Token{
				Type:  "number",
				Value: value,
			})

			continue
		}

		// if current is a double quotation,
		// means we're dealing with string
		if current == "\"" {
			// initializing an empty value store
			value := ""

			// skipping the current index,
			// as we don't care about the quotation
			i++

			// if next item is a blank space,
			// means it's an ending quotation we're dealing with
			// so we ignore it and move on
			if rgBlank.MatchString(string(input[i])) {
				continue
			}

			// checking each item after the quotation
			// if it's a valid string and not
			// another quotation
			for string(input[i]) != "\"" {
				// adding the current string item
				// in the value store
				value += string(input[i])

				// going for the next string item
				i++
			}

			// storing `value` in tokens array
			// as string type
			tokens = append(tokens, Token{
				Type:  "string",
				Value: value,
			})

			continue
		}

		rgAlpha := regexp.MustCompile(`[a-zA-Z]`) //regex for alphabets
		// if current is a alphabetical string,
		// means we're dealing with function names
		if rgAlpha.MatchString(current) {
			value := current

			// skipping current index as
			// it's already processed
			i++

			// we need to keep going until we
			// find a blank space
			for !rgBlank.MatchString(string(input[i])) {
				value += string(input[i])
				i++
			}

			// append `value` in tokens array as name
			tokens = append(tokens, Token{
				Type:  "name",
				Value: value,
			})

			continue
		}
	}

	return tokens
}
