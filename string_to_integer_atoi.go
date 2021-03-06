/*
https://leetcode.com/problems/string-to-integer-atoi/

Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).

The algorithm for myAtoi(string s) is as follows:

    Read in and ignore any leading whitespace.
    Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
    Read in next the characters until the next non-digit charcter or the end of the input is reached. The rest of the string is ignored.
    Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
    If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range. Specifically, integers less than -231 should be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
    Return the integer as the final result.

Note:

    Only the space character ' ' is considered a whitespace character.
    Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.
*/

package main

import "fmt"

func main() {
	fmt.Println(myAtoi("-2147483649"))
}

func myAtoi(s string) (integer int) {
	const maxDec = 1000000000
	const min = -2147483648
	const max = 2147483647
	const latestMinDigit = (-min) % 10
	const latestMaxDigit = max % 10

	i := 0
	for ; i < len(s) && s[i] == ' '; i++ {
	}

	minus := false
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		minus = s[i] == '-'
		i++
	}

	for ; i < len(s) && s[i] == '0'; i++ {
	}

	dec := maxDec
	overflow := false
	possibleOverflow := false
	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		digit := int(s[i] - '0')

		if dec == 0 || (dec == 1 && (overflow || (possibleOverflow && ((minus && digit > latestMinDigit) || (!minus && digit > latestMaxDigit))))) {
			if minus {
				return min
			}
			return max
		}

		if minus {
			integer -= digit * dec
		} else {
			integer += digit * dec
		}

		if dec == 10 {
			almostMax := max - latestMaxDigit
			almostMin := min + latestMinDigit
			overflow = (integer > almostMax) || (minus && integer < almostMin)
			possibleOverflow = (integer == almostMax) || (minus && integer == almostMin)
		}

		dec /= 10
	}

	if dec != 0 {
		integer /= (dec * 10)
	}

	return integer
}
