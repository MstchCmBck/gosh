package command

type redirection int

const (
	noredirection redirection = iota
	stdout
	stdoutappend
	stderr
	stderrappend
)
