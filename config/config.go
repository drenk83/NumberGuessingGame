package config

type GameMode int
type Language int

// Consts
// ---------------------------------
const (
	ClassicMode GameMode = iota
	TimedMode
)

const (
	LangEng Language = iota
	LangRus Language
)

const (
	DefaultMode        GameMode = ClassicMode
	DefaultMaxAttempts int      = 7
	DefaultLanguage    Language = LangEng
	DefaultMaxNumber   int      = 100
)

// Struct
// ---------------------------------
type GameConfig struct {
	mode        GameMode
	maxAttempts int
	language    Language
	maxNumber   int
}

// Constructors
// ---------------------------------
func NewGameConfig() *GameConfig {
	return &GameConfig{
		mode:        DefaultMode,
		maxAttempts: DefaultMaxAttempts,
		language:    DefaultLanguage,
		maxNumber:   DefaultMaxNumber,
	}
}

// Getters
// ---------------------------------
// func (g *GameConfig) Mode() string {
// 	return g.gameMode
// }

// func (g *GameConfig) MaxAtt() int {
// 	return g.maxAttempts
// }

// func (g *GameConfig) Lang() string {
// 	return g.language
// }

// func (g *GameConfig) Min() int {
// 	return g.min
// }

// func (g *GameConfig) Max() int {
// 	return g.max
// }

// Setters
// --------------------------------
// func (g *GameConfig) SetMode(m string) error {
// 	switch m {
// 	case ModeDefault:
// 		g.gameMode = ModeDefault
// 		return nil

// 	case ModeBot:
// 		g.gameMode = ModeBot
// 		return nil

// 	default:
// 		return errors.New("[ERROR] in changing the mode")
// 	}
// }

// func (g *GameConfig) SwapMode() {
// 	if g.gameMode == ModeDefault {
// 		g.gameMode = ModeBot
// 	} else {
// 		g.gameMode = ModeDefault
// 	}
// }

// func (g *GameConfig) SetLang(l string) error {
// 	switch l {
// 	case Eng:
// 		g.language = Eng
// 		return nil

// 	case Rus:
// 		g.language = Rus
// 		return nil

// 	default:
// 		return errors.New("[ERROR] in changing the language")
// 	}
// }

// func (g *GameConfig) SetAttempts(a int) error {
// 	if a <= 0 {
// 		return errors.New("[ERROR] Zero or negative number of attempts")
// 	}
// 	g.maxAttempts = a
// 	return nil
// }

// func (g *GameConfig) SetMinMax(mn, mx int) error {
// 	if mn > mx {
// 		return errors.New("[ERROR] Min is more than max")
// 	}
// 	if mn == mx {
// 		return errors.New("[ERROR] Same values")
// 	}
// 	g.min = mn
// 	g.max = mx
// 	return nil
// }

// Strings
// --------------------------------
func (g GameMode) String() string {
	switch g {
	case ClassicMode:
		return "Classic"
	case TimedMode:
		return "Timed"
	default:
		return "Unknown"
	}
}

func (l Language) String() string {
	switch l {
	case LangEng:
		return "English"
	case LangRus:
		return "Russian"
	default:
		return "Unknown"
	}
}
