


package main;

import "fmt";

// FIND DIGITS IN NUMBERS!
func findDigits(num int) int{
   var digits int = 0;
   for num != 0{
       num = num / 10;
	   digits++;
   }
   return digits;
}

func main(){

	val := findDigits(12345);
	fmt.Println(val); // OUTPUTS: 5!
}