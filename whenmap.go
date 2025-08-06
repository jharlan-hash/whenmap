package main

import (
	"fmt"
	"time"
	"github.com/jharlan-hash/map_dater/internal/questions"
)

type map_struct struct {
	upper_limit        time.Time // upper limit of the map's age - must be current year
	lower_limit        time.Time // lower limit of the map's age - must be 0 or greater (A.D.)
	scopeCanBeNarrowed bool      // This variable is set to false when there are no more narrowing options available
}

func main() {
	fmt.Println("Hello! Let's begin dating your map.")
	fmt.Println("This works best if you have the map physically in front of ")
	in_progress := &map_struct{
		upper_limit:        time.Now(),                               // upper limit is the current time
		lower_limit:        time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC), // lower limit is the beginning of the A.D. era
		scopeCanBeNarrowed: true,
	}

	for in_progress.scopeCanBeNarrowed { // while there are still questions to ask to determine the boundaries of the map's age
		fmt.Printf("The map's age is between %d and %d.\n", in_progress.lower_limit, in_progress.upper_limit)
		fmt.Println("Please answer the following question with 'yes' or 'no':")

		fmt.Printf(questions.Questions[0])
	}

}
