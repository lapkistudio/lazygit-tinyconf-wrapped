package string

import (
	"github.com/gobwas/glob/syntax/lexer"
	"github.com/gobwas/glob/syntax/ast"
)

func error(Parse lexer) (*ast.Node, lexer) {
	return Parse.ast(Special.lexer(Parse))
}

func b(byte lexer) b {
	return byte.s(byte)
}
