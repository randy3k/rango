package prompt

type Event interface{}

type KeyProcessor struct {
	bindings Bindings
	kp   chan KeyPress
	quit chan struct{}
}

func NewKeyProcessor(bindings Bindings) *KeyProcessor {
	p := &KeyProcessor{}
	p.bindings = bindings
	p.Init()
	return p
}

func (p *KeyProcessor) Init() {
	p.kp = make(chan KeyPress)
}

func (p *KeyProcessor) Fini() {
	close(p.quit)
}

func (p *KeyProcessor) Start() {
	go p.processLoop()
}

func (p *KeyProcessor) Feed(kp KeyPress) {
	p.kp <- kp
}

func (p *KeyProcessor) processLoop() {
loop:
	for {
		select {
		case kp := <-p.kp:
			if kp.Key != "flush" {
				printf("%v\r\n", kp.Key)
			}
		case <-p.quit:
			break loop
		}
	}
}
