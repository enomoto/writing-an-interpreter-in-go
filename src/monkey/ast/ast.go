package ast

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

