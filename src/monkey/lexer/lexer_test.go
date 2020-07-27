package lexer

import (
	"testing"
	"monkey/token"
)

func TestNextToken(t *testing.T) {
	tests := [] struct {
		exectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
	}
}

