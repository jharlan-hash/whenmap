package questions

import "time"

// Question represents a question to ask the user to narrow down the date of a map.
type Question struct {
	// QuestionText is the question posed to the user.
	QuestionText string

	// Earliest is the date when the subject of the question came into being.
	// If the user answers "yes", this becomes a potential new lower bound for the map's date.
	Earliest time.Time

	// Latest is the date when the subject of the question ceased to be.
	// If it still exists, this should be a time in the future (e.g., time.Now()).
	// If the user answers "yes", this becomes a potential new upper bound for the map's date.
	Latest time.Time

	// Importance helps prioritize questions. A higher number means a more important question.
	// Importance is often determined by how easy the question is to answer for a user.
	// For example, a question about a major country is more important than a small territorial change.
	// Range: 1 (least important) to 5 (most important).
	Importance int
}

// Questions is a list of all questions the program can ask.
// TODO: Expand this list to cover more regions and historical events.
var Questions = []Question{
	// --- AFRICA ---
	{
		QuestionText: "Does the country of South Sudan appear on the map?",
		Earliest:     time.Date(2011, 7, 9, 0, 0, 0, 0, time.UTC), // South Sudan independence
		Latest:       time.Now(),
		Importance:   5,
	},
	{
		QuestionText: "Does the country of Eritrea appear on the map?",
		Earliest:     time.Date(1993, 5, 24, 0, 0, 0, 0, time.UTC), // Eritrea independence
		Latest:       time.Now(),
		Importance:   5,
	},
	{
		QuestionText: "Does the country of Namibia appear on the map?",
		Earliest:     time.Date(1990, 3, 21, 0, 0, 0, 0, time.UTC), // Namibia independence
		Latest:       time.Now(),
		Importance:   5,
	},
	{
		QuestionText: "Is the country of Djibouti called 'Djibouti' (not 'French Somaliland' or 'Afars and Issas')?",
		Earliest:     time.Date(1977, 6, 27, 0, 0, 0, 0, time.UTC), // Djibouti independence
		Latest:       time.Now(),
		Importance:   3,
	},
	{
		QuestionText: "Are the Seychelles shown as an independent country (not a British colony)?",
		Earliest:     time.Date(1976, 6, 29, 0, 0, 0, 0, time.UTC), // Seychelles independence
		Latest:       time.Now(),
		Importance:   1,
	},
	{
		QuestionText: "Is the territory of Western Sahara called 'Western Sahara' (not 'Spanish Sahara')?",
		Earliest:     time.Date(1976, 2, 28, 0, 0, 0, 0, time.UTC), // Spain decolonized Spanish Sahara
		Latest:       time.Now(),
		Importance:   1,
	},
	{
		QuestionText: "Is Angola shown as an independent country (not a Portuguese colony)?",
		Earliest:     time.Date(1975, 11, 11, 0, 0, 0, 0, time.UTC), // Angola independence
		Latest:       time.Now(),
		Importance:   1,
	},
	{
		QuestionText: "Is Mozambique shown as an independent country (not a Portuguese colony)?",
		Earliest:     time.Date(1975, 6, 25, 0, 0, 0, 0, time.UTC), // Mozambique independence
		Latest:       time.Now(),
		Importance:   1,
	},
	{
		QuestionText: "Are Tanganyika and Zanzibar shown as a single country, 'Tanzania'?",
		Earliest:     time.Date(1964, 4, 26, 0, 0, 0, 0, time.UTC), // Unification of Tanganyika and Zanzibar
		Latest:       time.Now(),
		Importance:   1,
	},
	{
		QuestionText: "Is the area of modern-day Zimbabwe called 'Rhodesia'?",
		Earliest:     time.Date(1965, 11, 11, 0, 0, 0, 0, time.UTC), // Rhodesia's Unilateral Declaration of Independence
		Latest:       time.Date(1980, 4, 18, 0, 0, 0, 0, time.UTC),  // Became Zimbabwe
		Importance:   4,
	},
	{
		QuestionText: "Is the Democratic Republic of the Congo called 'Zaire'?",
		Earliest:     time.Date(1971, 10, 27, 0, 0, 0, 0, time.UTC), // Renamed to Zaire
		Latest:       time.Date(1997, 5, 17, 0, 0, 0, 0, time.UTC),  // Renamed back to DRC
		Importance:   3,
	},
}
