package lexer

import (
	"testing"
)

func TestTokenizeInteger(t *testing.T) {
	tokens, err := Tokenize("1")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(tokens) != 1 {
		t.Errorf("expected list of 1 token, got %d tokens", len(tokens))
	} else {
		if tokens[0].Type != INTEGER {
			t.Errorf("expected INTEGER token, got %s", tokens[0].Type.String())
		}
		if tokens[0].Literal != "1" {
			t.Errorf("expected token literal of '1', got '%s'", tokens[0].Literal)
		}
	}
}

func TestTokenizeFloat(t *testing.T) {
	tokens, err := Tokenize("100.9")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(tokens) != 1 {
		t.Errorf("expected list of 1 token, got %d tokens", len(tokens))
	} else {
		if tokens[0].Type != FLOAT {
			t.Errorf("expected FLOAT token, got %s", tokens[0].Type.String())
		}
		if tokens[0].Literal != "100.9" {
			t.Errorf("expected token literal of '100.9', got '%s'", tokens[0].Literal)
		}
	}
}

func checkToken(t *testing.T, token Token, tType TokenType, v string) {
	if token.Type != tType {
		t.Errorf("expected %s token, got %s", tType.String(), token.Type.String())
	}
	if token.Literal != v {
		t.Errorf("expected literal '%s', got '%s'", v, token.Literal)
	}
}

func TestTokenizeAdd(t *testing.T) {
	tokens, err := Tokenize("100.9 + 10")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(tokens) != 3 {
		t.Errorf("expected list of 3 token, got %d tokens", len(tokens))
	} else {
		checkToken(t, tokens[0], FLOAT, "100.9")
		checkToken(t, tokens[1], PLUS, "+")
		checkToken(t, tokens[2], INTEGER, "10")
	}
}

func TestTokenizeAddMinus(t *testing.T) {
	tokens, err := Tokenize("100.9 + 10 - 5")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(tokens) != 5 {
		t.Errorf("expected list of 5 token, got %d tokens", len(tokens))
	} else {
		checkToken(t, tokens[0], FLOAT, "100.9")
		checkToken(t, tokens[1], PLUS, "+")
		checkToken(t, tokens[2], INTEGER, "10")
		checkToken(t, tokens[3], MINUS, "-")
		checkToken(t, tokens[4], INTEGER, "5")
	}
}

func TestTokenizeAddNegative(t *testing.T) {
	tokens, err := Tokenize("100.9 -5")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(tokens) != 3 {
		t.Errorf("expected list of 3 token, got %d tokens", len(tokens))
	} else {
		checkToken(t, tokens[0], FLOAT, "100.9")
		checkToken(t, tokens[1], MINUS, "-")
		checkToken(t, tokens[2], INTEGER, "5")
	}
}

func TestTokenizeString(t *testing.T) {
	tokens, err := Tokenize(`"Hello world!"`)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(tokens) != 1 {
		t.Errorf("expected list of 1 token, got %d tokens", len(tokens))
	} else {
		checkToken(t, tokens[0], STRING, `"Hello world!"`)
	}
}

func TestTokenizeStringAddInteger(t *testing.T) {
	tokens, err := Tokenize(`"Hello world!" + 13`)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(tokens) != 3 {
		t.Errorf("expected list of 3 token, got %d tokens", len(tokens))
	} else {
		checkToken(t, tokens[0], STRING, `"Hello world!"`)
		checkToken(t, tokens[1], PLUS, "+")
		checkToken(t, tokens[2], INTEGER, "13")
	}
}
