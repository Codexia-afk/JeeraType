package generator

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
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

// LoadCustomWordlist validates custom wordlist file (≥50 words) and generates typing stream.
func LoadCustomWordlist(filePath string, targetDurationSec int) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("wordlist file not found: %s", filePath)
	}

	words := strings.Fields(string(content))
	if len(words) < 50 {
		return "", fmt.Errorf("wordlist must contain at least 50 words (found %d in %s)", len(words), filePath)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	targetWordCount := (targetDurationSec * 150) / 60
	if targetDurationSec <= 0 || targetWordCount < 60 {
		targetWordCount = 120
	}

	shuffled := make([]string, len(words))
	copy(shuffled, words)
	r.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	var resultWords []string
	idx := 0
	for len(resultWords) < targetWordCount {
		resultWords = append(resultWords, shuffled[idx%len(shuffled)])
		idx++
	}

	return strings.Join(resultWords, " "), nil
}

// GenerateText produces standard paragraph text with optional punctuation and numbers.
func GenerateText(targetDurationSec int, punctuation bool, numbers bool) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	targetWordCount := (targetDurationSec * 150) / 60
	if targetDurationSec <= 0 || targetWordCount < 60 {
		targetWordCount = 120
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
		for _, w := range words {
			cleanW := strings.Trim(w, ".,!?;:\"'")
			if !punctuation {
				cleanW = strings.ToLower(cleanW)
			} else {
				if r.Float32() < 0.15 {
					puncs := []string{",", ".", "!", "?", ";"}
					cleanW += puncs[r.Intn(len(puncs))]
				}
			}
			resultWords = append(resultWords, cleanW)

			if numbers && r.Float32() < 0.12 {
				numToken := fmt.Sprintf("%d", r.Intn(9999))
				resultWords = append(resultWords, numToken)
			}
		}
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
	if targetDurationSec <= 0 {
		targetChars = 1000
	}
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
func GenerateAdaptiveText(weakKeys []rune, weakBigrams []string, targetDurationSec int, punctuation bool, numbers bool) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	targetWordCount := (targetDurationSec * 150) / 60
	if targetDurationSec <= 0 || targetWordCount < 60 {
		targetWordCount = 120
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
		if !punctuation {
			w = strings.ToLower(strings.Trim(w, ".,!?;:\"'"))
		}
		resultWords = append(resultWords, w)
		if numbers && r.Float32() < 0.12 {
			resultWords = append(resultWords, fmt.Sprintf("%d", r.Intn(999)))
		}
	}

	return strings.Join(resultWords, " ")
}
