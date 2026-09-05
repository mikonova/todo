package tokens

const (
	LeftBrace           = '['
	RightBrace          = ']'
	ToDoMarker          = "TODO"
	UrgencyMarker       = "urgent"
	TagMarker           = "tag"
	LineStart           = '/'
	Star                = '*'
	Space               = ' '
	CompositeLineStart1 = "//"
	CompositeLineStart2 = "/*"
	CompositeLineEnd    = "*/"
	Colon               = ':'
	Comma               = ','
)

// TokenTypes
const (
	LineStartDeclaration = iota
	MultilineDeclaration
	MultilineUndeclaration
	LeftBraceIdentifier
	RightBraceIdentifier
	ToDoDeclaration
	ModifierDeclaration
	DescriptionDeclaration
)

// Modifiers
const (
	UrgencyModifier = iota
	TagModifier
	SubModifier
)

type Token struct {
	TokenType     int
	Modifier      int
	ModifierValue string
}
