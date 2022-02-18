package xno

type Player uint8

const (
	PlayerX Player = iota
	PlayerA
	PlayerC
	PlayerD
	PlayerM
	PlayerN
	PlayerO
	PlayerS
)

func (p Player) String() string {
	switch p {
	case PlayerX:
		return "X"
	case PlayerO:
		return "O"
	case PlayerA:
		return "👽"
	case PlayerC:
		return "😎"
	case PlayerD:
		return "🥸"
	case PlayerM:
		return "🧐"
	case PlayerN:
		return "🤓"
	case PlayerS:
		return "💀"
	default:
		return " "
	}
}
