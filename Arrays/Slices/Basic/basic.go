/* Slices
*  What Are These?
*  Unlike Arrays, Slices Are Flexible
*  a[lowest size : highest size]
 */

package main

import "fmt"

func basic() {
	primes := [5]int{2, 3, 5, 7, 11}

	var s = primes[1:4]
	fmt.Println(s)
}

func main() {
	basic()
}
