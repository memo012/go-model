package template

type LocalDoc struct {
	DocSuper
}

func NewLocalDoc() *LocalDoc {
	l := new(LocalDoc)
	l.DocSuper.GetContent = l.GetContent
	return l
}

func (l *LocalDoc) GetContent() string {
	return "this is a LocalDoc"
}
