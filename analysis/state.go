package analysis

import (
	"cc-lsp/lsp"
	"fmt"
	"regexp"
	"strings"
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
	prefixes := []string{"feat", "fix", "chore", "test", "docs"}
	middleRegex := ""
	for _, item := range prefixes {
		middleRegex += regexp.QuoteMeta(item) + "|"
	}

	// Build the regex pattern dynamically based on the list of prefixes
	pattern := `^(?:` + middleRegex + `)(?:\(.+\))?:\s+`

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

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Characters: %d", uri, len(document)),
		},
	}
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {

	// Ask your static analysis tools to figure out good completions
	items := []lsp.CompletionItem{
		{
			Label:         "Neovim (BTW)",
			Detail:        "Very cool editor",
			Documentation: "Fun to watch in videos. Don't forget to like & subscribe to streamers using it :)",
		},
	}

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
