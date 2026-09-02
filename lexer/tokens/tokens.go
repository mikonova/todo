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

const (
	LeftBraceIdentifier = iota
	RightBraceIdentifier
	ToDoDeclaration
	Modifier
	DescriptionDeclaration
)

const (
	UrgencyModifier = iota
	tagModifier
)

type Token struct {
	TokenType     int
	Modifier      int
	ModifierValue string
}
