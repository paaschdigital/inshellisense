package model

import "github.com/google/uuid"

type Subcommand struct {
	Name           []string //single or array string, required
	Description    string
	Args           []Arg
	Options        []Option
	Subcommands    []Subcommand
	FilterStrategy FilterStrategy
}

type Option struct {
	Name         []string //single or array string, required
	Args         []Arg    //single or array Arg, optional
	Description  string   //single, optional
	IsPersistent bool
	ExclusiveOn  []string
}

type Arg struct {
	Name           string //single, optional
	Description    string //single, optional
	Templates      []Template
	Suggestions    []Suggestion
	FilterStrategy FilterStrategy
	Generator      *Generator
	IsVariadic     bool
	IsOptional     bool
	IsCommand      bool
}

type Suggestion struct {
	Name        []string
	Description string
}

type TermSuggestions struct {
	ArgumentDescription string
	Suggestions         []TermSuggestion
}

type TermSuggestion struct {
	Name        string
	Description string
	Type        TermSuggestionType
}

type TermSuggestionType string

const (
	TermSuggestionTypeFolder     TermSuggestionType = "folder"
	TermSuggestionTypeFile       TermSuggestionType = "file"
	TermSuggestionTypeArg        TermSuggestionType = "arg"
	TermSuggestionTypeSubcommand TermSuggestionType = "subcommand"
	TermSuggestionTypeOption     TermSuggestionType = "option"
	TermSuggestionTypeAI         TermSuggestionType = "ai"
	TermSuggestionTypeDefault    TermSuggestionType = ""
)

type ProcessedToken struct {
	Token   string
	Persist bool
}

type Generator struct {
	Id          uuid.UUID
	Script      string
	Function    func([]string) []TermSuggestion
	PostProcess func(string) []TermSuggestion
	Template    []Template
	SplitOn     string
	SkipCache   bool
}

type Template string

const (
	TemplateFilepaths Template = "filepaths"
	TemplateFolders   Template = "folders"
	TemplateHistory   Template = "history"
	TemplateHelp      Template = "help"
)

type FilterStrategy string

const (
	FilterStrategyPrefix FilterStrategy = "prefix"
	FilterStrategyFuzzy  FilterStrategy = "fuzzy"
	FilterStrategyEmpty  FilterStrategy = ""
)

var (
	Templates = []Template{TemplateFilepaths, TemplateFolders, TemplateHistory, TemplateHelp}
	TermIcons = map[TermSuggestionType]string{
		TermSuggestionTypeFolder:     "📁",
		TermSuggestionTypeFile:       "📄",
		TermSuggestionTypeSubcommand: "📦",
		TermSuggestionTypeOption:     "💲",
		TermSuggestionTypeArg:        "💪",
		TermSuggestionTypeDefault:    "💪",
		TermSuggestionTypeAI:         "🔮",
	}
)