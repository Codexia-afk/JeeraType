package generator

import (
	"math/rand"
	"strings"
	"time"
)

var goSnippets = []string{
	"package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, JeeraType!\")\n}",
	"type Server struct {\n    addr string\n    port int\n}\n\nfunc NewServer(a string, p int) *Server {\n    return &Server{addr: a, port: p}\n}",
	"func worker(id int, jobs <-chan int, results chan<- int) {\n    for j := range jobs {\n        results <- j * 2\n    }\n}",
	"if err != nil {\n    log.Fatalf(\"fatal error: %v\", err)\n}",
}

var pythonSnippets = []string{
	"def calculate_wpm(keystrokes, elapsed_sec):\n    minutes = elapsed_sec / 60.0\n    return (keystrokes / 5) / minutes if minutes > 0 else 0",
	"class DatabaseConnection:\n    def __init__(self, db_path):\n        self.db_path = db_path\n        self.conn = None",
	"import json\nwith open('data.json', 'r') as f:\n    data = json.load(f)\nprint(data)",
	"async def fetch_payload(url):\n    async with aiohttp.ClientSession() as session:\n        async with session.get(url) as response:\n            return await response.json()",
}

var jsSnippets = []string{
	"const calculateAccuracy = (correct, total) => {\n    if (total === 0) return 100;\n    return (correct / total) * 100;\n};",
	"async function loadDataset() {\n    const res = await fetch('/api/words');\n    const data = await res.json();\n    return data.words;\n}",
	"document.addEventListener('keydown', (e) => {\n    if (e.key === 'Escape') {\n        resetTest();\n    }\n});",
	"export class EventEmitter {\n    constructor() {\n        this.listeners = new Map();\n    }\n}",
}

// GenerateLanguageCodeText returns language-specific code snippets.
func GenerateLanguageCodeText(lang string, targetDurationSec int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var bank []string

	switch strings.ToLower(lang) {
	case "python", "py":
		bank = pythonSnippets
	case "javascript", "js":
		bank = jsSnippets
	case "go", "golang":
		bank = goSnippets
	default:
		bank = goSnippets
	}

	shuffled := make([]string, len(bank))
	copy(shuffled, bank)
	r.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	var snippets []string
	totalChars := 0
	targetChars := targetDurationSec * 8
	if targetDurationSec <= 0 {
		targetChars = 500
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
