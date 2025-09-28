package utils

import (
	"fmt"
	"strings"
)

// Base.

// comment splits a string into lines and comments each line.
// Returns commented lines joined back into a single string.
func comment(desc string) (comment string) {
	descs := strings.Split(desc, "\n")
	comments := make([]string, 0, len(descs))
	// Find the last line.
	var last int
	for last = len(descs) - 1; last >= 0; last-- {
		if d := descs[last]; len(d) > 0 {
			break
		}
	}
	// Comment each line.
	for i := 0; i <= last; i++ {
		d := strings.TrimRight(descs[i], asciiSpaces)
		c := fmt.Sprintf("// " + d)
		comments = append(comments, c)
	}
	comment = strings.Join(comments, "\n")
	return
}

// Cap capitalizes the first character of a string.
func Cap(name string) (capped string) {
	capped = strings.ToUpper(name[:1]) + name[1:]
	return
}

// DeCap un capitalizes the first character of a string.
func DeCap(name string) (decapped string) {
	decapped = strings.ToLower(name[:1]) + name[1:]
	return
}

// LabelToVarName converts a label to a valid varName
// Ex: "helllo world." to "helloWorld"
func LabelToVarName(label string) (varName string) {
	name := LabelToName(label)
	// Lower case camel.
	varName = strings.ToLower(name[:1]) + name[1:]
	return
}

// LabelToName converts a label to a valid name.
// Ex: "helllo world." to "HelloWorld"
// err is nil or contains the error message for the user.
func LabelToName(label string) (name string) {
	validNameChars := make([]string, 0, len(label))
	var started bool
	var followingSpace bool
	for i := range label {
		ch := label[i : i+1]
		switch {
		case strings.Contains(asciiSpaces, ch):
			followingSpace = true
			// No white spaces and punctuation allowed.
		case !started:
			// The first character must be a capital letter.
			switch {
			case strings.Contains(ucAlphabet, ch):
				started = true
				followingSpace = false
				validNameChars = append(validNameChars, ch)
			case strings.Contains(lcAlphabet, ch):
				started = true
				followingSpace = false
				validNameChars = append(validNameChars, strings.ToUpper(ch))
			}
		case started:
			// Only characters and numbers follow the first capitalized letter.
			switch {
			case strings.Contains(ucAlphabet, ch):
				validNameChars = append(validNameChars, ch)
				followingSpace = false
			case strings.Contains(lcAlphabet, ch):
				if followingSpace {
					// Enforce camel case.
					ch = strings.ToUpper(ch)
					followingSpace = false
				}
				validNameChars = append(validNameChars, ch)
			case strings.Contains(digits, ch):
				validNameChars = append(validNameChars, ch)
				followingSpace = false
			}
		}
	}
	name = strings.Join(validNameChars, "")
	return
}

// validateMessageName returns if the message name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
func validateMessageName(name string) (isValid bool, userMessage string) {
	isValid, userMessage = validateCamelCaseName(name, "message")
	return
}

// validateRecordName returns if the record name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
func validateRecordName(name string) (isValid bool, userMessage string) {
	isValid, userMessage = validateCamelCaseName(name, "record")
	return
}

// validateScreenPanelName returns if the screen panel name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
func validateScreenPanelName(name string) (isValid bool, userMessage string) {
	isValid, userMessage = validateCamelCaseName(name, "panel")
	return
}

// validateScreenTabName returns if the screen tab name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
func validateScreenTabName(name string) (isValid bool, userMessage string) {
	isValid, userMessage = validateCamelCaseName(name, "tab")
	return
}

// validateScreenAccordionItemName returns if the screen accordionItem name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
func validateScreenAccordionItemName(name string) (isValid bool, userMessage string) {
	isValid, userMessage = validateCamelCaseName(name, "accordionItem")
	return
}

// validateScreenName returns if the screen name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
func validateScreenName(name string) (isValid bool, userMessage string) {
	isValid, userMessage = validateCamelCaseName(name, "screen")
	return
}

// validateCamelCaseName returns if the name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
// Param nameType is what the name is created for. ("record", "panel")
func validateCamelCaseName(name, nameType string) (isValid bool, userMessage string) {

	defer func() {
		if !isValid {
			userMessage = fmt.Sprintf("The %s name %q must be TitleCase.", nameType, name)
		}
	}()

	// First char must be upper case.
	s := name[:1]
	if s != strings.ToUpper(s) {
		isValid = false
		return
	}
	// Letters only
	isValid = false
	for _, c := range name {
		isValid = strings.ContainsRune(ucAlphabet, c)
		if !isValid {
			isValid = strings.ContainsRune(lcAlphabet, c)
		}
		if !isValid {
			return
		}
	}
	return
}

// Prefix returns a prefixed version of a string.
func Prefix(s string, prefix string) (prefixed string) {
	prefixed = prefix + s
	return
}

// Suffix returns a suffixed version of a string.
func Suffix(s string, suffix string) (suffixed string) {
	suffixed = s + suffix
	return
}

// Prefix returns a prefixed and suffixed version of a string.
func PrefixSuffix(s string, prefix, suffix string) (fixed string) {
	fixed = prefix + s + suffix
	return
}

// Suffix combinations.

func SuffixLowerCase(s string, suffix string) (lowerCased string) {
	suffixed := Suffix(s, suffix)
	lowerCased = strings.ToLower(suffixed)
	return
}

// CapSuffix returns a suffixed and then capped version of a string.
func CapSuffix(s string, suffix string) (capped string) {
	suffixed := Suffix(s, suffix)
	capped = Cap(suffixed)
	return
}

// DeCapSuffix returns a suffixed and then decapped version of a string.
func DeCapSuffix(s string, suffix string) (decapped string) {
	suffixed := Suffix(s, suffix)
	decapped = DeCap(suffixed)
	return
}

// Prefix combinations.

// PrefixLowerCase returns a prefixed and lowercased version of a string.
func PrefixLowerCase(s string, prefix string) (lowerCased string) {
	prefixed := Prefix(s, prefix)
	lowerCased = strings.ToLower(prefixed)
	return
}

// PrefixCap returns a prefixed and capped version of a string.
func PrefixCap(s string, prefix string) (capped string) {
	prefixed := Prefix(s, prefix)
	capped = Cap(prefixed)
	return
}

// PrefixDeCap returns a prefixed and decapped version of a string.
func PrefixDeCap(s string, prefix string) (decapped string) {
	prefixed := Prefix(s, prefix)
	decapped = DeCap(prefixed)
	return
}

// Prefix suffix combinations.

// PrefixCapSuffix returns a prefixed, capped and suffixed version of a string.
func PrefixCapSuffix(s string, prefix, suffix string) (capped string) {
	fixed := PrefixSuffix(s, prefix, suffix)
	capped = Cap(fixed)
	return
}

// PrefixCapSuffix returns a prefixed, decapped and suffixed version of a string.
func PrefixDeCapSuffix(s string, prefix, suffix string) (decapped string) {
	fixed := PrefixSuffix(s, prefix, suffix)
	decapped = DeCap(fixed)
	return
}
