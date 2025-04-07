package token

import (
	_ "embed"
	"encoding/json"
)

//go:embed tokens.json
var tokensJSON []byte

type Tokens struct {
	Map map[string]*Token
}

type Token struct {
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	ID       int    `json:"id"`
	Decimals int    `json:"decimals"`
}

func NewTokens() *Tokens {
	t := &Tokens{
		Map: make(map[string]*Token),
	}

	var list []*Token
	if err := json.Unmarshal(tokensJSON, &list); err != nil {
		return nil
	}

	for _, token := range list {
		t.Map[token.Address] = token
		t.Map[token.Symbol] = token
	}
	return t
}
