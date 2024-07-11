package titlescreen

type Button struct {
	text      string
	isClicked bool
	onClick   func()
}

func NewButton(txt string, onclick func()) Button {
	return Button{
		text:      txt,
		isClicked: false,
		onClick:   onclick,
	}
}

func (btn Button) Run() {
	btn.onClick()
}
