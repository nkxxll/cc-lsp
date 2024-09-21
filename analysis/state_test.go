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
		{"feat(main): add new feature", false},      // Starts with a conventional prefix
		{"fix: correct a bug", false},               // Starts with a conventional prefix
		{"test(stats): clean up codebase", false},   // Starts with a conventional prefix
		{"test(stats)!: clean up codebase", false},  // Starts with a conventional prefix
		{"test!: clean up codebase", false},         // Starts with a conventional prefix
		{"!test: update documentation", true},       // Does not start with a conventional prefix
		{"test!(this): update documentation", true}, // Does not start with a conventional prefix
		{"doggoo: update documentation", true},      // Does not start with a conventional prefix
		{"doggoo!: update documentation", true},     // Does not start with a conventional prefix
		{"blueberry: change structure", true},       // Does not start with a conventional prefix
		{"yeet: add unit tests", true},              // Does not start with a conventional prefix
		{"this is any message", true},               // Does not start with a conventional prefix
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

	if len(diagnostics) != 7 {
		t.Fatalf("diagnostics should be 7 long is %d long", len(diagnostics))
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
		{`


`, false, ""},
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

func TestGetWord(t *testing.T) {
	cases := []struct {
		line     string
		position lsp.Position
		expected string
	}{
		{"this is a nice text", lsp.Position{Line: 4, Character: 4}, " "},
		{"this is a nice text", lsp.Position{Line: 4, Character: 5}, "is"},
		{"this is a nice text", lsp.Position{Line: 4, Character: 8}, "a"},
		{"this is a nice text", lsp.Position{Line: 4, Character: 10}, "nice"},
		{"this is a nice text", lsp.Position{Line: 4, Character: 0}, "this"},
		{"this is a nice text", lsp.Position{Line: 4, Character: 1}, "this"},
		{"this is a nice text", lsp.Position{Line: 4, Character: 2}, "this"},
		{"this is a nice text", lsp.Position{Line: 4, Character: 3}, "this"},
		{"this is a :nice: text", lsp.Position{Line: 4, Character: 10}, ":"},
		{"this is a :nice: text", lsp.Position{Line: 4, Character: 11}, "nice"},
		{"this is a nice: text", lsp.Position{Line: 4, Character: 11}, "nice"},
		{"this :is a nice: text", lsp.Position{Line: 4, Character: 6}, "is"},
		{"this :is a nice: text", lsp.Position{Line: 4, Character: 7}, "is"},
		{"a b c d e f g", lsp.Position{Line: 4, Character: 2}, "b"},
		{"a b c d e f g", lsp.Position{Line: 4, Character: 0}, "a"},
		{"", lsp.Position{Line: 4, Character: 3}, ""},
	}

	for idx, tc := range cases {
		word := getWord(tc.line, tc.position)
		if word != tc.expected {
			t.Fatalf("Test case %d failed. Got %s - Exp %s", idx, word, tc.expected)
		}
	}
}
