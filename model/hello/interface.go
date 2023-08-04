package hello

type Hello struct {
	Msg string `json:"msg"`
}

func (h *Hello) GetHelloSeeme() {
	h.Msg = "Hello Seeme! []~(￣▽￣)~*"
}
