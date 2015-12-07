package main

import "bytes"

const (
	Indent         = "  "
	NewLine        = "\n"
	SymbolSpace    = " "
	V_ARGUMENT_SEP = ": "
)

//
// CodeEmitter.
//

type CodeEmitter struct {
	Buffer bytes.Buffer

	indentLevel int
}

func (e *CodeEmitter) Dedent() {
	e.indentLevel -= 1
	if e.indentLevel < 0 {
		e.indentLevel = 0
	}
}

func (e *CodeEmitter) Indent() {
	e.indentLevel += 1
}

func (e *CodeEmitter) Write(strings ...string) {
	for i := 0; i < e.indentLevel; i++ {
		e.Buffer.WriteString(Indent)
	}

	// so for example, if given {"module", "Stripe"}, this would print
	// "module<SPACE>Stripe"
	for i, s := range strings {
		if i > 0 {
			e.Buffer.WriteString(SymbolSpace)
		}

		e.Buffer.WriteString(s)
	}

	// trailing newline
	e.Buffer.WriteString(NewLine)
}

//
// Emit definitions.
//

func (n *ArgumentListNode) Emit(e *CodeEmitter) {
	strings := []string{S_ARG_L}

	for _, a := range n.RequiredArguments {
		strings = append(strings, a.Symbol, ",")
	}

	for _, a := range n.OptionalArguments {
		strings = append(strings, a.Symbol, V_ARGUMENT_SEP, ",")
	}

	for _, a := range n.DefaultArguments {
		strings = append(strings, a.Symbol.Symbol, V_ARGUMENT_SEP, a.DefaultValue.ScalarValue(), ",")
	}

	strings = append(strings, S_ARG_R)
	e.Write(strings...)
}

func (n *BeginNode) Emit(e *CodeEmitter) {
	e.Write(S_BEGIN)
	emitStatementListBody(e, n.Body)
}

func (n *ClassNode) Emit(e *CodeEmitter) {
	// leading newline
	e.Write("")
	e.Write(S_CLASS, n.Symbol.Symbol)
	emitStatementListBody(e, n.Body)
}

func (n *CommentNode) Emit(e *CodeEmitter) {
	e.Write(S_COMMENT, n.Comment)
}

func (n *DefaultArgumentNode) Emit(e *CodeEmitter) {
	panic("unexpected call")
}

func (n *MethodNode) Emit(e *CodeEmitter) {
	s := emitToString(n.Arguments)
	e.Write(S_DEF, n.Symbol.Symbol, s)
	emitStatementListBody(e, n.Body)
}

func (n *ModuleNode) Emit(e *CodeEmitter) {
	// leading newline
	e.Write("")
	e.Write(S_MODULE, n.Symbol.Symbol)
	emitStatementListBody(e, n.Body)
}

func (n *RequireNode) Emit(e *CodeEmitter) {
	e.Write(S_REQUIRE, n.File.ScalarValue())
}

func (n *StatementListNode) Emit(e *CodeEmitter) {
	for _, node := range n.Children {
		node.Emit(e)
	}
}

func (n *StringScalarNode) Emit(e *CodeEmitter) {
	panic("unexpected call")
}

func (n *SymbolNode) Emit(e *CodeEmitter) {
	panic("unexpected call")
}

func emitToString(n Node) string {
	emitter := &CodeEmitter{}
	n.Emit(emitter)
	return emitter.Buffer.String()
}

func emitStatementListBody(e *CodeEmitter, n *StatementListNode) {
	e.Indent()
	n.Emit(e)
	e.Dedent()
	e.Write(S_END)
	e.Write("")
}
