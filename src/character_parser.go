package src

type CharacterKind int

const (
	CharNumber CharacterKind = iota
	CharColon
	CharEmpty
)

type Character struct {
	kind  CharacterKind
	value int
}

func parseTimeToCharacters(timeText string) []Character {
	characters := make([]Character, 0, len(timeText))

	for _, r := range timeText {
		switch {
		case r >= '0' && r <= '9':
			characters = append(characters, Character{
				kind:  CharNumber,
				value: int(r - '0'),
			})
		case r == ':':
			characters = append(characters, Character{
				kind: CharColon,
			})
		default:
			characters = append(characters, Character{
				kind: CharEmpty,
			})
		}
	}

	return characters
}
