package bookanalyser

import (
	"fmt"
	"regexp"
	"strings"
)

type Book struct {
	length        int
	uniquqWords   int
	uniqueQuota   float32
	dialogueQuota float32
}

func (b *Book) Read() {
	raw := "“William Ann!” she said, turning, “how are you son she said?”"
	b.Analyse(raw)
}

//Analyse is the main function
func (b *Book) Analyse(text string) {
	sampleRegexp := regexp.MustCompile(`[^A-Za-z_ ]`)
	lower := strings.ToLower(text)
	clean := sampleRegexp.ReplaceAllString(lower, "")

	words := strings.Split(clean, " ")
	length := len(words)

	uniques := make(map[string]bool)
	for i := range words {
		uniques[words[i]] = true
	}

	sampleRegexp = regexp.MustCompile(`(“.*?”)`)
	dialogues := sampleRegexp.FindAllString(text, -1)
	countOfDialogue := 0
	for i := range dialogues {
		countOfDialogue += len(strings.Split(dialogues[i], " "))
	}

	b.dialogueQuota = float32(countOfDialogue) / float32(length)
	b.uniquqWords = len(uniques)
	b.length = length
	b.uniqueQuota = float32(b.uniquqWords) / float32(b.length)

}

func (b *Book) Print() {
	fmt.Printf("%+v\n", b)
}
