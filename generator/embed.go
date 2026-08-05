package generator

import (
	_ "embed"
	"encoding/json"
	"math/rand"
	"strings"
	"time"
)

//go:embed paragraphs.json
var embeddedParagraphsData []byte

//go:embed code_snippets.json
var embeddedCodeData []byte

//go:embed quotes.json
var embeddedQuotesData []byte

type paragraphDataset struct {
	Paragraphs []string `json:"paragraphs"`
}

type codeDataset struct {
	CodeSnippets []string `json:"code_snippets"`
}

type quoteDataset struct {
	Quotes []string `json:"quotes"`
}

var (
	cachedParagraphs []string
	cachedCode       []string
	cachedQuotes     []string
)

func init() {
	var pData paragraphDataset
	if err := json.Unmarshal(embeddedParagraphsData, &pData); err == nil && len(pData.Paragraphs) > 0 {
		cachedParagraphs = pData.Paragraphs
	} else {
		cachedParagraphs = []string{
			"The quick brown fox jumps over the lazy dog. Terminal applications provide rapid feedback and high efficiency.",
		}
	}

	var cData codeDataset
	if err := json.Unmarshal(embeddedCodeData, &cData); err == nil && len(cData.CodeSnippets) > 0 {
		cachedCode = cData.CodeSnippets
	} else {
		cachedCode = []string{
			"func main() {\n    fmt.Println(\"JeeraType Code Mode\")\n}",
		}
	}

	var qData quoteDataset
	if err := json.Unmarshal(embeddedQuotesData, &qData); err == nil && len(qData.Quotes) > 0 {
		cachedQuotes = qData.Quotes
	} else {
		cachedQuotes = []string{
			"Simplicity is prerequisite for reliability. Complex systems tend to fail in mysterious ways.",
		}
	}
}

// GenerateText produces standard paragraph text.
func GenerateText(targetDurationSec int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	targetWordCount := (targetDurationSec * 150) / 60
	if targetWordCount < 60 {
		targetWordCount = 60
	}

	shuffled := make([]string, len(cachedParagraphs))
	copy(shuffled, cachedParagraphs)
	r.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	var resultWords []string
	idx := 0
	for len(resultWords) < targetWordCount {
		para := shuffled[idx%len(shuffled)]
		words := strings.Fields(para)
		resultWords = append(resultWords, words...)
		idx++
	}

	return strings.Join(resultWords, " ")
}

// GenerateCodeText returns code snippets joined into a typing passage.
func GenerateCodeText(targetDurationSec int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	shuffled := make([]string, len(cachedCode))
	copy(shuffled, cachedCode)
	r.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	var snippets []string
	totalChars := 0
	targetChars := targetDurationSec * 8
	for _, snip := range shuffled {
		snippets = append(snippets, snip)
		totalChars += len(snip)
		if totalChars >= targetChars {
			break
		}
	}
	return strings.Join(snippets, "\n\n")
}

// GenerateQuoteText returns quotes joined into a typing passage.
func GenerateQuoteText() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return cachedQuotes[r.Intn(len(cachedQuotes))]
}

// GenerateAdaptiveText generates text heavy in the user's weak keys and bigram transitions.
func GenerateAdaptiveText(weakKeys []rune, weakBigrams []string, targetDurationSec int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	targetWordCount := (targetDurationSec * 150) / 60
	if targetWordCount < 60 {
		targetWordCount = 60
	}

	var allWords []string
	for _, p := range cachedParagraphs {
		words := strings.Fields(p)
		allWords = append(allWords, words...)
	}

	var prioritizedWords []string
	for _, word := range allWords {
		lowerWord := strings.ToLower(word)
		containsWeak := false

		for _, k := range weakKeys {
			if strings.ContainsRune(lowerWord, k) {
				containsWeak = true
				break
			}
		}
		if !containsWeak {
			for _, bg := range weakBigrams {
				if strings.Contains(lowerWord, bg) {
					containsWeak = true
					break
				}
			}
		}
		if containsWeak {
			prioritizedWords = append(prioritizedWords, word)
		}
	}

	if len(prioritizedWords) == 0 {
		prioritizedWords = allWords
	}

	var resultWords []string
	for len(resultWords) < targetWordCount {
		w := prioritizedWords[r.Intn(len(prioritizedWords))]
		resultWords = append(resultWords, w)
	}

	return strings.Join(resultWords, " ")
}
