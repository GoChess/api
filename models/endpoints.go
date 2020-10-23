package models

import "time"

// Ping is used to return a true value upon pining
type Ping struct {
	Pong bool `json: "pong"`
}

// Preview used for short previews of different blog posts
type Preview struct {
	Title  string `json: "title"`
	FirstP string `json: "firstP"`
	Date   time.Time
	Author string `json: "author"`
}

// BlogParagraph provides paragraphs, imagery and
type BlogParagraph struct {
	Image           string `json: "url", omitempty`
	ImageAnnotation string `json: "annotation", omitempty`
	Paragraph       string `json: "paragraph", omitempty`
}

// Blog superstruct of other structs to allow blog creation
type Blog struct {
	BlogPreview Preview
	Paragraphs  []BlogParagraph
}
