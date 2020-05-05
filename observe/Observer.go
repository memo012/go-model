package observe

type Observer interface {
	Update(subject *Subject)
}
