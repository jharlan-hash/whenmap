package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/jharlan-hash/whenmap/internal/questions"
)

// MapDateRange represents the possible date range for a map.
type MapDateRange struct {
	Earliest time.Time // The lower bound of the map's possible date.
	Latest   time.Time // The upper bound of the map's possible date.
}

func main() {
	fmt.Println("Hello! Let's begin dating your map.")
	fmt.Println("This works best if you have the map physically in front of you.")
	fmt.Println("Please answer the following questions with 'yes', 'no', or 'idk'.")

	dateRange := &MapDateRange{
		Latest:   time.Now(),
		Earliest: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	// TODO: Implement a better question selection strategy based on importance
	// and the current date range.
	for _, q := range questions.Questions {
		// Check if the question is relevant to the current date range.
		// A question is relevant if its date range overlaps with the map's possible date range.
		if dateRange.Earliest.Before(q.Latest) && q.Earliest.Before(dateRange.Latest) {
			fmt.Printf("\nBased on your answers, the map was created between %s and %s.\n",
				dateRange.Earliest.Format("January 2, 2006"),
				dateRange.Latest.Format("January 2, 2006"))
			askQuestion(q, dateRange)
		}
	}

	fmt.Printf("\nFinished! The final estimated date range for your map is between %s and %s.\n",
		dateRange.Earliest.Format("January 2, 2006"),
		dateRange.Latest.Format("January 2, 2006"))
}

// askQuestion poses a question to the user and updates the date range based on the answer.
func askQuestion(q questions.Question, dateRange *MapDateRange) {
	fmt.Println(q.QuestionText)

	var answer string
	_, err := fmt.Scanln(&answer)
	if err != nil {
		// If Scanln returns an error, it's likely EOF, so we can just exit.
		return
	}

	switch strings.ToLower(strings.TrimSpace(answer)) {
	case "yes":
		// The answer is "yes", so the map's date must be within the question's date range.
		// We narrow the map's date range to the intersection of the current range and the question's range.
		dateRange.Earliest = maxTime(dateRange.Earliest, q.Earliest)
		dateRange.Latest = minTime(dateRange.Latest, q.Latest)
	case "no":
		// The answer is "no". This is more complex.
		// For a question like "Does South Sudan exist?", a "no" implies the map was made
		// before South Sudan's independence.
		// For a question about an event with a start and end date, like "Does Rhodesia exist?",
		// a "no" could mean before OR after the event. The current logic handles the "before"
		// case, but not the "after" case, which would require supporting multiple date ranges.
		// We are simplifying here and assuming "no" means "before the event started".
		dateRange.Latest = minTime(dateRange.Latest, q.Earliest)
	case "idk":
		// User doesn't know. We don't learn anything, so we don't update the range.
		fmt.Println("Okay, we'll skip that one.")
	default:
		fmt.Println("Invalid input. Please answer 'yes', 'no', or 'idk'.")
		askQuestion(q, dateRange) // Ask the same question again.
	}
}

// maxTime returns the later of two times.
func maxTime(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}

// minTime returns the earlier of two times.
func minTime(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}
