package analysis

import (
	"cc-lsp/lsp"
	"regexp"
	"strings"
	"unicode"
)

type State struct {
	// Map of file names to contents
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

// diagnoseNoConventionalCommitMsg evaluates the first line that has text in
// the git commit and returns an diagnose if the line does not match a
// conventional commit format
func diagnoseNoConventionalCommitMsg(text string) (lsp.Diagnostic, bool) {
	// todo: expand the list to the angular conventional commits
	middleRegex := ""
	for _, item := range lsp.Prefixes {
		middleRegex += regexp.QuoteMeta(item) + "|"
	}

	// Build the regex pattern dynamically based on the list of prefixes
	pattern := `^(?:` + middleRegex + `)(?:\(.+\))?!?:\s+`

	// Compile the regex
	re := regexp.MustCompile(pattern)
	if re.FindStringIndex(text) == nil {
		diagnostic := lsp.Diagnostic{
			Range:    LineRange(0, 0, 0),
			Severity: 1,
			Source:   "cc-lint",
			Message:  "First line should start with the type of the commit in a conventional commit. (e.g. feat, fix, ...)",
		}
		return diagnostic, true
	}
	return lsp.Diagnostic{}, false
}

// getFirstLine gets the first line from the git commit that is not empty or a comment
func getFirstLine(text string) (string, bool) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		return line, true
	}
	return "", false
}

func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	// todo: do we want to lint trailing white space?
	diagnostics := []lsp.Diagnostic{}
	firstLine, ok := getFirstLine(text)

	if ok {
		// see if the line starts with a conventional commit type like: feat, fix, ...
		diagnose, found := diagnoseNoConventionalCommitMsg(firstLine)
		if found {
			diagnostics = append(diagnostics, diagnose)
		}
	}

	return diagnostics
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text)
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	// In real life, this would look up the type in our type analysis code...

	document := s.Documents[uri]
	line := strings.Split(document, "\n")[position.Line]
	word := getWord(line, position)
	content, ok := lsp.HoverContents[word]
	if ok {
		return lsp.HoverResponse{
			Response: lsp.Response{
				RPC: "2.0",
				ID:  &id,
			},
			Result: lsp.HoverResult{
				Contents: string(content),
			},
		}

	}
	// for now return empty hover
	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: "No Information",
		},
	}
}

func getWord(line string, position lsp.Position) string {
	// if line is empty there is no word to find
	if line == "" {
		return ""
	}
	character := position.Character
	res := string(line[character])
	// check whether we hit a space
	if res == " " || !unicode.IsLetter(rune(res[0])) {
		return res
	}

	// check whether we are out of bounds
	if !(character+1 >= len(line)) {

		// go to the end of the word
		for unicode.IsLetter(rune(line[character+1])) {
			character += 1
			res += string(line[character])
			// check if we are at the end of the line in the next run
			if character+1 >= len(line) {
				break
			}
		}
	}
	// go the other way reset the character position
	character = position.Character
	// check for bound again
	if !(character-1 < 0) {
		// go to the start of the word
		for unicode.IsLetter(rune(line[character-1])) {
			character -= 1
			res = string(line[character]) + res
			// check if we are at the start of the line in the next run
			if character-1 < 0 {
				break
			}
		}
	}
	return res
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {

	// Ask your static analysis tools to figure out good completions
	items := lsp.GetCompletions()
	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}

	return response
}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}
