package memento

type InputText struct {
	content string
}

func (i *InputText) Append(content string) {
	i.content += content
}

func (i *InputText) GetText() string {
	return i.content
}

func (i *InputText) SnapShot() *Snapshot {
	return &Snapshot{content: i.content}
}

func (i *InputText) Restore(snapshot *Snapshot) {
	i.content = snapshot.GetText()
}

type Snapshot struct {
	content string
}

func (s *Snapshot) GetText() string {
	return s.content
}
