package analysis

import (
	"cc-lsp/lsp"
	"testing"
)

func TestConventionalCommitPrefix(t *testing.T) {
	// Define test cases
	testCases := []struct {
		line     string
		expected bool
	}{
		{"feat(main): add new feature", false}, // Starts with a conventional prefix
		{"fix: correct a bug", false},          // Starts with a conventional prefix
		{"chore: clean up codebase", false},    // Starts with a conventional prefix
		{"docs: update documentation", true},   // Does not start with a conventional prefix
		{"refactor: change structure", true},   // Does not start with a conventional prefix
		{"test: add unit tests", true},         // Does not start with a conventional prefix
	}

	var diagnostics []lsp.Diagnostic
	for _, tc := range testCases {
		// Test if the line does NOT match the conventional commit prefix
		diagnose, found := diagnoseNoConventionalCommitMsg(tc.line)
		if found != tc.expected {
			t.Fatal("this should be an error but is not")
		}
		if found {
			diagnostics = append(diagnostics, diagnose)
		}
	}

	if len(diagnostics) != 3 {
		t.Fatalf("diagnostics should be 3 long is %d long", len(diagnostics))
	}

	for _, item := range diagnostics {
		if item.Range.Start.Line != 0 || item.Range.End.Line != 0 {
			t.Fatalf("Line: The problem should be on the start of the line not start %d end %d", item.Range.Start.Character, item.Range.End.Character)
		}
		if item.Range.Start.Character != 0 || item.Range.End.Character != 0 {
			t.Fatalf("Character: The problem should be on the start of the line not start %d end %d", item.Range.Start.Character, item.Range.End.Character)
		}
	}
}

func TestGetFirstLine(t *testing.T) {
	texts := []struct {
		text string
		ok   bool
		line string
	}{{`# this is the best thing I have ever done
# lorem ipsum
# what is love
# baby don't hurt me
feat: new commit`, true, "feat: new commit"},
		{"", false, ""},
		{"# this is only a comment", false, ""},
		{"# feat: this is also only a comment", false, ""},
		{"feat: this is also only a right line with a # comment", true, "feat: this is also only a right line with a # comment"},
		{`# this is a comment

feat: this is also only a right line with a # comment`, true, "feat: this is also only a right line with a # comment"},
	}
	for idx, tc := range texts {
		expected := "feat: new commit"
		first, ok := getFirstLine(tc.text)
		if ok != tc.ok {
			t.Fatalf("Failed at iteration %d; There should be a line found", idx)
		}
		if first != tc.line {
			t.Fatalf("Failed at iteration %d; first line should be this %s", idx, expected)
		}
	}
}
