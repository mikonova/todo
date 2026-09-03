package parser

import (
	"errors"
	"os"
	"todo/errdef"
	"todo/lexer"
	"todo/lexer/tokens"
)

// Take tokens and parse them into AST
func ParseAll(storage []lexer.TokenStorage) error {
	parsedTokenStream := make([]tokens.Token, 0)
	newToken := tokens.Token{}
	for _, sample := range storage {
		file, err := os.OpenFile(lexer.StorageFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
		if err != nil {
			return errors.New(errdef.ErrBase + "Cannot open the summary file")
		}
		file.WriteString("[" + sample.FileName + "]\n")
		for idx, token := range sample.TokenStream {
			newToken = recognizeToken(token, idx, sample.TokenStream)

			parsedTokenStream = append(parsedTokenStream, newToken)
		}
	}
}

func recognizeToken(token string, index int, s []string) tokens.Token {
	newToken := tokens.Token{
		TokenType:     -1,
		Modifier:      -1,
		ModifierValue: "",
	}
	switch token {
	case tokens.CompositeLineStart1:
		newToken.TokenType = tokens.LineStartDeclaration
	case tokens.CompositeLineStart2:
		newToken.TokenType = tokens.LineStartDeclaration
	case tokens.CompositeLineEnd:
		newToken.TokenType = tokens.MultilineUndeclaration
	case tokens.ToDoMarker:
		newToken.TokenType = tokens.ToDoDeclaration
	case string(tokens.LeftBrace):
		newToken.TokenType = tokens.LeftBraceIdentifier
	case string(tokens.RightBrace):
		newToken.TokenType = tokens.RightBraceIdentifier
	case tokens.TagMarker:
		newToken.TokenType = tokens.ModifierDeclaration
		newToken.Modifier = tokens.TagModifier
	case tokens.UrgencyMarker:
		newToken.TokenType = tokens.ModifierDeclaration
		newToken.Modifier = tokens.UrgencyModifier
	default:
		if s[index-1] == string(tokens.RightBrace) {
			newToken.TokenType = tokens.DescriptionDeclaration
			newToken.Modifier = tokens.SubModifier
			newToken.ModifierValue = token
		}
		if s[index-1] == string(tokens.Colon) &&
	}
	return newToken
}

/*
подписанные токены -> поток токенов для парсера -> поток именованных токенов
*/
