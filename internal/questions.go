package questions
import "time"

type question struct {
	questionText string
	earliest     time.Time // If the question is answered "yes", this is the current lower bound of the map's age
	latest       time.Time // If the question is answered "yes", this is the current upper bound of the map's age
	importance   int       // Importance of the question, used to prioritize questions. Mostly determined by ease of answering
	// e.g. if the question asks the user to find a tiny detail on the map, it is less important than a question asking about a major event
	// Range: 1 (least important) to 5 (most important)
}

var Questions = []question{
	// AFRICAN SECTION
	{
		questionText: "Does South Sudan exist?",
		earliest:     time.Date(2011, 7, 9, 0, 0, 0, 0, time.UTC), // South Sudan became independent on July 9, 2011
		latest:       time.Now(),
		importance:   5,
	},
	{
		questionText: "Does Eritrea exist?",
		earliest:     time.Date(1993, 5, 24, 0, 0, 0, 0, time.UTC), // Eritrea became independent on May 24, 1993
		latest:       time.Now(),
		importance:   5,
	},
	{
		questionText: "Does Namibia exist",
		earliest:     time.Date(1990, 3, 21, 0, 0, 0, 0, time.UTC), // Namibia became independent on March 21, 1990
		latest:       time.Now(),
		importance:   5,
	},
	{
		questionText: "Is it called Djibouti?",
		earliest:     time.Date(1977, 6, 27, 0, 0, 0, 0, time.UTC), // Djibouti became independent on June 27, 1977
		latest:       time.Now(),
		importance:   3,
	},
	{
		questionText: "Are the Seychelles still a British colony?",
		earliest:     time.Date(1976, 6, 29, 0, 0, 0, 0, time.UTC), // Seychelles became independent on June 29, 1976
		latest:       time.Now(),
		importance:   1,
	},
	{
		questionText: "Has Spanish Sahara been renamed Western Sahara?",
		earliest:     time.Date(1976, 2, 28, 0, 0, 0, 0, time.UTC), // Western Sahara was declared a Spanish colony until February 28, 1976
		latest:       time.Now(),
		importance:   1,
	},
	{
		questionText: "Is Angola still a Portuguese colony?",
		earliest:     time.Date(1975, 11, 11, 0, 0, 0, 0, time.UTC), // Angola became independent on November 11, 1975
		latest:       time.Now(),
		importance:   1,
	},
	{
		questionText: "Is Mozambique still a Portuguese colony?",
		earliest:     time.Date(1975, 6, 25, 0, 0, 0, 0, time.UTC), // Mozambique became independent on June 25, 1975
		latest:       time.Now(),
		importance:   1,
	},
	{
		questionText: "Does Tanzania include Zanzibar?",
		earliest:     time.Date(1964, 4, 26, 0, 0, 0, 0, time.UTC), // Zanzibar was merged with Tanganyika to form Tanzania on April 26, 1964
		latest:       time.Now(),
		importance:   1,
	},
	{
		questionText: "Does Rhodesia exist?",
		earliest:     time.Date(1965, 11, 11, 0, 0, 0, 0, time.UTC), // Rhodesia was declared a self-governing colony on November 11, 1965
		latest:       time.Date(1980, 4, 18, 0, 0, 0, 0, time.UTC),  // Rhodesia became Zimbabwe on April 18, 1980
		importance:   4,
	},
	{
		questionText: "Is it called Zaire or the Democratic Republic of the Congo?",
		earliest:     time.Date(1971, 5, 17, 0, 0, 0, 0, time.UTC), // Zaire was renamed to the Democratic Republic of the Congo on May 17, 1997
		latest:       time.Now(),
	},
}
