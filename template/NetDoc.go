package template

type NetDoc struct {
	DocSuper
}

func NewNetDoc() *NetDoc {
	c := new(NetDoc)
	c.DocSuper.GetContent = c.GetContent
	return c
}

func (n *NetDoc) GetContent() string {
	return "this is a NetDoc"
}


