package lsp

type Content string

// the content of the over response for the different message types
const (
	buildContent    Content = "build: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)"
	ciContent       Content = "ci: Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)"
	docsContent     Content = "docs: Documentation only changes"
	featContent     Content = "feat: A new feature"
	fixContent      Content = "fix: A bug fix"
	perfContent     Content = "perf: A code change that improves performance"
	refactorContent Content = "refactor: A code change that neither fixes a bug nor adds a feature"
	styleContent    Content = "style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"
	testContent     Content = "test: Adding missing tests or correcting existing tests"
)

// maps message types to hover content
var HoverContents = map[string]Content{
	"build":    buildContent,
	"ci":       ciContent,
	"docs":     docsContent,
	"feat":     featContent,
	"fix":      fixContent,
	"perf":     perfContent,
	"refactor": refactorContent,
	"style":    styleContent,
	"test":     testContent,
}

var Prefixes = []string{
	"build",
	"ci",
	"docs",
	"feat",
	"fix",
	"perf",
	"refactor",
	"style",
	"test",
}

func GetCompletions() []CompletionItem {
	completions := []CompletionItem{}
	const keyword = 14
	for _, item := range Prefixes {
		documentation, ok := HoverContents[item]
		if !ok {
			panic("There is no documentation for the given keyword!")
		}
		completion := CompletionItem{
			Label:         item,
			Kind:          keyword,
			Detail:        string(documentation),
			Documentation: string(documentation),
		}
		completions = append(completions, completion)
	}
	return completions
}
