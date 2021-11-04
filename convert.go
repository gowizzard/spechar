//********************************************************************************************************************//
//
// This file is part of spechar.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package spechar

import "strings"

// Convert is to convert the string & remove all special characters
func Convert(text string) string {

	// Create list for special characters
	list := []string{"^", "!", "\"", "§", "$", "%", "&", "/", "(", ")", "=", "ß", "°", "[", "]", "|", "{", "}", "?", "+", "*", "#", "'", "<", ">", ",", ".", "-", ";", ":", "_"}

	// Check each special characters & replace ti with nothing
	for _, value := range list {
		text = strings.Replace(text, value, "", -1)
	}

	// Return text without umlauts
	return text

}
