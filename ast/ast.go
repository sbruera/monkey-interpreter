package ast

import "monkey/token"

type Node interface {
  TokenLiteral() string
}

type Statment interface {
  Node
  statmentNode()
}

type Expresion interface {
  Node 
  expressionNode()
}

type Program struct {
  Statments[]Statment
}

func (p *Program) TokenLiteral() string {
  if len(p.Statments) > 0 {
    return p.Statments[0].TokenLiteral()
  } else {
    return "";
  }
}

type LetStatment struct {
  Token token.Token //the token.LET token
  Name *Identifier
  Value Expresion
}

func (ls *LetStatment) statmentNode() {}
func (ls *LetStatment) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
  Token token.Token // The token.IDENT token
  Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

