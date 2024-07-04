package main

type Wisdom struct {
	Description string `json:"description"`
	Explanation string `json:"explanation"`
}

var BaseWisdomList []Wisdom = []Wisdom{
	{Description: "Have fun", Explanation: "Having fun in work is a big deal when it comes to the creative process that is software engineering (or engineering in general). There will always be tasks that are less fun. Performing those tasks first will make you look forward to the fun tasks."},
	{Description: "Don't Repeat Yourself", Explanation: "Duplication leads to maintenance hell and logic contradictions, reducing seems a Good Thing™."},
}
