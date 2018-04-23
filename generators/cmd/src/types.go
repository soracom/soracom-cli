package main

type stringFlag struct {
	VarName      string
	LongOption   string
	DefaultValue string
	ShortHelp    string
	In           string
	Name         string
	Required     bool
}

type stringFlagsByName []stringFlag

func (s stringFlagsByName) Len() int {
	return len(s)
}

func (s stringFlagsByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s stringFlagsByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

type integerFlag struct {
	VarName      string
	LongOption   string
	DefaultValue int64
	Format       string
	ShortHelp    string
	In           string
	Name         string
	Required     bool
}

type integerFlagsByName []integerFlag

func (s integerFlagsByName) Len() int {
	return len(s)
}

func (s integerFlagsByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s integerFlagsByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

type floatFlag struct {
	VarName      string
	LongOption   string
	DefaultValue float64
	Format       string
	ShortHelp    string
	In           string
	Name         string
	Required     bool
}

type floatFlagsByName []floatFlag

func (s floatFlagsByName) Len() int {
	return len(s)
}

func (s floatFlagsByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s floatFlagsByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

type boolFlag struct {
	VarName      string
	LongOption   string
	DefaultValue bool
	Format       string
	ShortHelp    string
	In           string
	Name         string
	Required     bool
}

type boolFlagsByName []boolFlag

func (s boolFlagsByName) Len() int {
	return len(s)
}

func (s boolFlagsByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s boolFlagsByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

type commandArgs struct {
	Endpoint                  string
	Use                       string
	Short                     string
	Long                      string
	CommandVariableName       string
	ParentCommandVariableName string
	FileName                  string
	RequireAuth               bool
	RequireOperatorID         bool
	BodyExists                bool
	Method                    string
	BasePath                  string
	Path                      string
	ContentType               string
	StringFlags               []stringFlag
	StringSliceFlags          []stringFlag
	IntegerFlags              []integerFlag
	FloatFlags                []floatFlag
	BoolFlags                 []boolFlag
}
