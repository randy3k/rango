package prompt

type BindingHandler func([]Key, []rune)

type Binding struct {
	Keys    []Key
	Handler BindingHandler
	Ctx     []Context
	Eager   bool
}

// TODO: cache context
func (b *Binding) Enabled() bool {
	for _, v := range b.Ctx {
		if !v() {
			return false
		}
	}
	return true
}

func (b *Binding) Match(keys []Key) bool {
	if len(keys) != len(b.Keys) {
		return false
	}
	for i, v := range b.Keys {
		if v != keys[i] && v != Any {
			return false
		}
	}
	return true
}

func (b *Binding) PreMatch(keys []Key) bool {
	if len(keys) >= len(b.Keys) {
		return false
	}
	for i, v := range keys {
		if v != b.Keys[i] {
			return false
		}
	}
	return true
}

func (b *Binding) Copy() *Binding {
	return &Binding{
		Keys: append([]Key{}, b.Keys...),
		Handler: b.Handler,
		Ctx: append([]Context{}, b.Ctx...),
		Eager: b.Eager,
	}
}

func (b *Binding) Normalize() *Binding {
	newb := b.Copy()
	for i, k := range newb.Keys {
		if newk, ok := KeyAliasMap[k]; ok {
			newb.Keys[i] = newk
		}
	}
	return newb
}


type KeyBindings struct {
	Bindings []*Binding
}

func (kbs *KeyBindings) Normalize() *KeyBindings{
	bindings := make([]*Binding, len(kbs.Bindings))
	newkbs := &KeyBindings{Bindings: bindings}
	for i, b := range kbs.Bindings {
		bindings[i] = b.Normalize()
	}
	return newkbs
}


func (kbs *KeyBindings) Prepend(bindings ...*Binding) {
	kbs.Bindings = append(bindings, kbs.Bindings...)
}

func (kbs *KeyBindings) Append(bindings ...*Binding) {
	kbs.Bindings = append(kbs.Bindings, bindings...)
}

func (kbs *KeyBindings) HasPrefix(keys []Key) bool {
	for _, bind := range kbs.Bindings {
		if bind.PreMatch(keys) && bind.Enabled() {
			return true
		}
	}
	return false
}

func (kbs *KeyBindings) Get(keys []Key, earger bool) (*Binding, bool) {
	for _, bind := range kbs.Bindings {
		if bind.Eager == earger && bind.Match(keys) && bind.Enabled() {
			return bind, true
		}
	}

	return nil, false
}
