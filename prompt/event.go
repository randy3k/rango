package prompt

type Event struct {
	Keys []Key
	Data []rune
	Prompt *Prompt
}
