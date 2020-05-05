package observe

type Subject struct {
	obs []Observer  // 群聊
	context string // 更新群聊信息
}

func NewSubject () *Subject{
	return &Subject{obs: make([]Observer, 0)}
}

func (s *Subject) notify()  {
	for _, o := range s.obs { // 每个人都通知一次
		o.Update(s) // 更新通知
	}
}

func (s *Subject) UpdateContext(context string)  {
	s.context = context
	s.notify() // 更新后通知
}

func (s *Subject) Attach(o Observer)  {
	s.obs = append(s.obs, o)
}
