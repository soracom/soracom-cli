package main

// Pagination holds information about pagination
type Pagination struct {
	Response Response `yaml:"response"`
	Request  Request  `yaml:"request"`
}

// Response holds information about response
type Response struct {
	Header string `yaml:"header"`
}

// Request holds information about request
type Request struct {
	Param string `yaml:"param"`
}

type stringFlag struct {
	VarName                string
	LongOption             string
	DefaultValueSpecified  bool
	DefaultValue           string
	ShortHelp              string
	In                     string
	Name                   string
	Required               bool
	HarvestFilesPathEscape bool
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
	VarName               string
	LongOption            string
	DefaultValueSpecified bool
	DefaultValue          int64
	Format                string
	ShortHelp             string
	In                    string
	Name                  string
	Required              bool
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
	VarName               string
	LongOption            string
	DefaultValueSpecified bool
	DefaultValue          float64
	Format                string
	ShortHelp             string
	In                    string
	Name                  string
	Required              bool
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
	VarName               string
	LongOption            string
	DefaultValueSpecified bool
	DefaultValue          bool
	Format                string
	ShortHelp             string
	In                    string
	Name                  string
	Required              bool
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
	Endpoint                          string
	Use                               string
	Short                             string
	Long                              string
	CommandVariableName               string
	ParentCommandVariableName         string
	FileName                          string
	RequireAuth                       bool
	RequireOperatorID                 bool
	BodyExists                        bool
	SendBodyRaw                       bool
	ResponseBodyRaw                   bool
	Method                            string
	BasePath                          string
	Path                              string
	ContentType                       string
	ContentTypeFromArg                bool
	ContentTypeVarName                string
	PathParamsExist                   bool
	QueryParamsExist                  bool
	StringFlags                       []stringFlag
	StringSliceFlags                  []stringFlag
	IntegerFlags                      []integerFlag
	FloatFlags                        []floatFlag
	BoolFlags                         []boolFlag
	RequiredFlagExists                bool
	PaginationAvailable               bool
	PaginationKeyHeaderInResponse     string
	PaginationRequestParameterInQuery string
	Deprecated                        bool
	AlternativeCommand                string
	HasArrayResponse                  bool
	Tag                               string
	OperationID                       string
}
