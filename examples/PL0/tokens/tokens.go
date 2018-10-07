package tokens

type Token int

const (
	Period     Token = iota // .
	Const                   // const
	Comma                   // ,
	Semicolon               // ;
	Var                     // var
	Procedure               // procedure
	Assignment              // :=
	Exclam                  // !
	Ques                    // ?
	Call                    // call
	Begin                   // begin
	End                     // end
	If                      // if
	Then                    // then
	While                   // while
	Do                      // do
	Odd                     // odd
	Equal                   // =
	Hash                    // #
	LT                      // <
	LTE                     // <=
	GT                      // >
	GTE                     // >=
	Plus                    // +
	Minus                   // -
	Mul                     // *
	Div                     // /
	OpenParen               // (
	CloseParen              // )
)

var (
	all = []Token{Period, Const, Comma, Semicolon, Var, Procedure, Assignment, Exclam, Ques, Call, Begin, End, If,
		Then, While, Do, Odd, Equal, Hash, LT, LTE, GT, GTE, Plus, Minus, Mul, Div, OpenParen, CloseParen}
	m map[string]Token
)

func init() {
	m = make(map[string]Token)
	for _, token := range all {
		m[token.String()] = token
	}
}

func TokenFromString(s string) (Token, bool) {
	token, ok := m[s]
	return token, ok
}
