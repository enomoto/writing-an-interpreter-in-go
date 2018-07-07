package ast

import "monkey/token"

// Interface

type Node interface {
  TokenLiteral() string
}

type Statement interface {
  Node
  statementNode()
}

type Experession interface {
  Node
  expressionNode()
}

// Implementation

type Program struct {
  Statements []Statement
}

func (p *Program) TokenLiteral() string {
  if len (p.Statements) > 0 {
    return p.Statements[0].TokenLiteral()
  } else {
    return ""
  }
}

type LetStatement struct {
  Token token.Token // token.LET トークン
  Name *Identifier
  Value Experession
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
  Token token.Token // token.IDENT トークン
  Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// return 文 (return <expression>;) を表現する構造体
type ReturnStatement struct {
  Token token.Token // token.RETURN トークン
  ReturnValue Experession
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
