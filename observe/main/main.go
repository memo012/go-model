package main

import "designMode/observe"

func main() {
	subject := observe.NewSubject()
	r1 := observe.NewReader("memo011")
	r2 := observe.NewReader("memo012")
	r3 := observe.NewReader("memo013")
	subject.Attach(r1)
	subject.Attach(r2)
	subject.Attach(r3)
	subject.UpdateContext("有人来来...")
}
