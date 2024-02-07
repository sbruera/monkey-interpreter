package ast

import (
	"bytes"
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// Stament don't produce values
type Statment interface {
	Node
	statmentNode()
}

// Expressions produce some value
type Expresion interface {
	Node
	expressionNode()
}

type Program struct {
	Statments []Statment
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statments {
		out.WriteString(s.String())
	}
	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statments) > 0 {
		return p.Statments[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatment struct {
	Token token.Token //the token.LET token
	Name  *Identifier //the variable name
	Value Expresion   //the expresion after the variable name
}

func (ls *LetStatment) statmentNode()        {}
func (ls *LetStatment) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatment) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

type ReturnStament struct {
	Token       token.Token //The return token
	ReturnValue Expresion
}

func (rs *ReturnStament) statmentNode()        {}
func (rs *ReturnStament) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStament) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token     token.Token //The first token of our expression
	Expresion Expresion
}

func (es *ExpressionStatement) statmentNode()        {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expresion != nil {
		return es.Expresion.String()
	}
	return ""
}

type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}
