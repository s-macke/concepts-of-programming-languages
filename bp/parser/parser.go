package parser

import (
	"fmt"
	"github.com/s-macke/concepts-of-programming-languages/bp/ast"
	"io"
)

type Tokenizer interface {
	NextToken() (string, error)
}

type Parser struct {
	Tokenizer Tokenizer
	rootNode  ast.Node
	token     string // ll(1)
}

// 	<expression> ::= <term> { <or> <term> }
// 	<term> ::= <factor> { <and> <factor> }
// 	<factor> ::= <var> | <not> <factor> | (<expression>)
// 	<or>  ::= '|'
// 	<and> ::= '&'
// 	<not> ::= '!'
//  <var> ::= '[a-zA-Z0-9]+'

func (p *Parser) Parse() (ast.Node, error) {
	p.rootNode = nil
	if err := p.expression(); err != nil {
		return nil, err
	}
	return p.rootNode, nil
}

func (p *Parser) expression() error {
	if err := p.term(); err != nil {
		return err
	}
	for p.token == "|" {
		lhs := p.rootNode
		if err := p.term(); err != nil {
			return err
		}
		rhs := p.rootNode
		p.rootNode = &ast.Or{LeftExpr: lhs, RightExpr: rhs}
	}
	return nil
}

func (p *Parser) term() error {
	if err := p.factor(); err != nil {
		return err
	}
	for p.token == "&" {
		lhs := p.rootNode
		if err := p.factor(); err != nil {
			return err
		}
		rhs := p.rootNode
		p.rootNode = &ast.And{LeftExpr: lhs, RightExpr: rhs}
	}
	return nil
}

func (p *Parser) factor() error {
	token, err := p.Tokenizer.NextToken()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	p.token = token
	if p.token == "!" {
		if err := p.factor(); err != nil {
			return err
		}
		p.rootNode = &ast.Not{Expr: p.rootNode}
	} else if p.token == "(" {
		if err := p.expression(); err != nil {
			return err
		}
		p.token, err = p.Tokenizer.NextToken()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	} else {
		p.rootNode = ast.Var{Name: p.token}
		p.token, err = p.Tokenizer.NextToken()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Parser) String() string {
	return fmt.Sprint(p.rootNode)
}
