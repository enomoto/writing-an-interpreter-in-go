package token

type TokenType string

const (
  ILLEGAL = "ILLEGAL"
  EOF     = "EOF"

  // 識別子 + リテラル
  IDENT = "IDENT"
  INT = "INT"

  // 演算子
  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"

  LT = "<"
  GT = ">"

  // デリミタ
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // キーワード
  FUNCTION = "FUNCTION"
  LET = "LET"
  TRUE = "TRUE"
  FALSE = "FALSE"
  IF = "IF"
  RETURN = "RETURN"
  ELSE = "ELSE"
)

type Token struct {
  Type    TokenType
  Literal string
}

var keywords = map[string]TokenType{
  "fn": FUNCTION,
  "let": LET,
  "true": TRUE,
  "false": FALSE,
  "if": IF,
  "return": RETURN,
  "else": ELSE,
}

func LookupIdent(ident string) TokenType {
  if tok, ok:= keywords[ident]; ok {
    return tok
  }
  return IDENT
}
