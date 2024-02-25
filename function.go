package main

// func concat(s1 string, s2 string) string {
// 	return s1 + s2
// }

// func main() {
// 	fmt.Println(concat("Lane,", " happy bday!"))
// 	fmt.Println(concat("Elon,", " hope that!"))
// 	fmt.Println(concat("Go,", " is great!"))
// }

func yearsUntilEvents(age int) (
	yearsUntilAdults,
	yearsUntilDrinking,
	yearsUntilCarRental int,
) {
	yearsUntilAdults = 18 - age
	if yearsUntilAdults < 0 {
		yearsUntilAdults = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdults, yearsUntilDrinking, yearsUntilCarRental
}
