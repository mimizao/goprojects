package workerpool2

type Option func(*Pool)

func WithBlock(block bool) Option {
	return func(p *Pool) {
		p.block = block
	}
}

func WirhPreAllocWorkers(preAlloc bool) Option {
	return func(p *Pool) {
		p.preAlloc = preAlloc
	}
}
