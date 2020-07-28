package key

type KeyBinding struct {
	Keys    []Key
	// we use interface{} because we don't want to fix the function signatures
	Handler interface{}
	Context     []Context
	Eager   bool
}

// TODO: cache context
func (kb *KeyBinding) Enabled() bool {
	for _, v := range kb.Context {
		if !v() {
			return false
		}
	}
	return true
}

func (kb *KeyBinding) Match(keys []Key) bool {
	if len(keys) != len(kb.Keys) {
		return false
	}
	for i, v := range kb.Keys {
		if v != keys[i] && v != Any {
			return false
		}
	}
	return true
}

func (kb *KeyBinding) PreMatch(keys []Key) bool {
	if len(keys) >= len(kb.Keys) {
		return false
	}
	for i, v := range keys {
		if v != kb.Keys[i] {
			return false
		}
	}
	return true
}

func (kb *KeyBinding) Copy() *KeyBinding {
	return &KeyBinding{
		Keys: append([]Key{}, kb.Keys...),
		Handler: kb.Handler,
		Context: append([]Context{}, kb.Context...),
		Eager: kb.Eager,
	}
}

func (kb *KeyBinding) Normalize() *KeyBinding {
	newkb := kb.Copy()
	for i, k := range newkb.Keys {
		if newk, ok := keyAliasMap[k]; ok {
			newkb.Keys[i] = newk
		}
	}
	return newkb
}

type KeyBindings []*KeyBinding


func (kbs *KeyBindings) Normalize() *KeyBindings{
	bindings := make(KeyBindings, len(*kbs))
	for i, b := range *kbs {
		bindings[i] = b.Normalize()
	}
	return &bindings
}


func (kbs *KeyBindings) Prepend(bindings ...*KeyBinding) {
	*kbs = append(bindings, *kbs...)
}

func (kbs *KeyBindings) Append(bindings ...*KeyBinding) {
	*kbs = append(*kbs, bindings...)
}

func (kbs *KeyBindings) HasPrefix(keys []Key) bool {
	for _, bind := range *kbs {
		if bind.PreMatch(keys) && bind.Enabled() {
			return true
		}
	}
	return false
}

func (kbs *KeyBindings) Get(keys []Key, earger bool) (*KeyBinding, bool) {
	for _, bind := range *kbs {
		if bind.Eager == earger && bind.Match(keys) && bind.Enabled() {
			return bind, true
		}
	}

	return nil, false
}
