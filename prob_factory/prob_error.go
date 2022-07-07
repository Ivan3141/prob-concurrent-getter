package prob_factory

type SyntaxError struct {
	Msg string
}

func (e SyntaxError) Error() string { return "syntax error: " + e.Msg }
