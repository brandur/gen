package main

const (
	S_ARG_L   = "("
	S_ARG_R   = ")"
	S_BEGIN   = "begin"
	S_CLASS   = "class"
	S_COMMENT = "#"
	S_DEF     = "def"
	S_END     = "end"
	S_MODULE  = "module"
	S_REQUIRE = "require"
)

type (
	Node interface {
		Emit(e *CodeEmitter)
		Name() string
	}

	ScalarNode interface {
		ScalarValue() string
	}

	ArgumentListNode struct {
		RequiredArguments []*SymbolNode
		OptionalArguments []*SymbolNode
		DefaultArguments  []*DefaultArgumentNode
	}

	BeginNode struct {
		Body *StatementListNode
	}

	CommentNode struct {
		Comment string
	}

	ClassNode struct {
		Symbol *SymbolNode
		Body   *StatementListNode
	}

	DefaultArgumentNode struct {
		Symbol *SymbolNode
		DefaultValue ScalarNode
	}

	MethodNode struct {
		Symbol *SymbolNode
		Arguments *ArgumentListNode
		Body   *StatementListNode
	}

	ModuleNode struct {
		Symbol *SymbolNode
		Body   *StatementListNode
	}

	RequireNode struct {
		File *StringScalarNode
	}

	StatementListNode struct {
		Children []Node
	}

	StringScalarNode struct {
		Value string
	}

	SymbolNode struct {
		Symbol string
	}
)

func (n *ArgumentListNode) Name() string  { return "(...)" }
func (n *BeginNode) Name() string         { return S_BEGIN }
func (n *ClassNode) Name() string         { return S_CLASS }
func (n *CommentNode) Name() string       { return S_COMMENT + " ..." }
func (n *MethodNode) Name() string        { return S_DEF }
func (n *ModuleNode) Name() string        { return S_MODULE }
func (n *RequireNode) Name() string       { return S_REQUIRE }
func (n *StatementListNode) Name() string { return "" }
func (n *StringScalarNode) Name() string  { return n.Value }
func (n *SymbolNode) Name() string        { return n.Symbol }

//
// scalars
//

func (n *StringScalarNode) ScalarValue() string {
	return "\"" + n.Value + "\""
}
