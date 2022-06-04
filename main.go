/*
Given a string s, count the number/frequency of each
characters. The frequency of a character is the number of times it appears in the given
string.

Example 1:

Input: s = "book"
Output: [b:1 k:1 o:2]

Example 2:

Input: s = "zwzwzw"
Output: [w:3 z:3]


Example 3:
Input: s = "mississippi"
Output: [i:4 m:1 p:2 s:4]
*/

package main

import "fmt"

func charCount(st string) map[string]int {

	m := map[string]int{}

	for _, s := range st {

		m[string(s)]++

	}

	return m
}

func main() {

	fmt.Println(charCount("book"))
	fmt.Println(charCount("zwzwzw"))
	fmt.Println(charCount("mississippi"))
}
