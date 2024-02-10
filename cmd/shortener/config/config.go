package config

type Args struct {
	Start_addr string
	Base_addr  string
}

type GetArgsBuilder interface {
	SetStart(string) GetArgsBuilder
	SetBase(string) GetArgsBuilder
	Build() *Args
}
type ConcreteGetArgsBuilder struct {
	args *Args
}

func NewGetArgsBuilder() *ConcreteGetArgsBuilder {
	return &ConcreteGetArgsBuilder{args: &Args{}}
}

func (cgab *ConcreteGetArgsBuilder) SetStart(start_addr string) GetArgsBuilder {
	cgab.args.Start_addr = start_addr
	return cgab
}

func (cgab *ConcreteGetArgsBuilder) SetBase(base_addr string) GetArgsBuilder {
	cgab.args.Base_addr = base_addr
	return cgab
}
func (cgab *ConcreteGetArgsBuilder) Build() *Args {
	return cgab.args
}
