package prompt


type Binding struct {
	Keys []Key
	Handler func(Event)
	Context [](func() bool)
}


// TODO: cache context
func (b *Binding) Enabled() bool {
	for _, v := range b.Context {
		if !v() {
			return false
		}
	}
	return true
}

func (b *Binding) Dispatch(event Event) {
	b.Handler(event)
}


type Bindings struct {
	Binds []Binding
}


func (b *Bindings) Add(keys []Key, h func(Event), context ...func() bool) {
	b.Binds = append(b.Binds, Binding{Keys: keys, Handler: h, Context: context})
}
