package position

// Position represents a chess position.
// A position can be used to generate a list of legal moves or validate a move is legal.
type Position struct {
	Pieces [12]uint64

	WhitesTurn     bool
	CastlingRights [4]bool
	EnPassant      uint8
}

// New creates a new starting position.
func New() Position {
	return Position{

		Pieces: [12]uint64{
			0xFF_00,                   // White pawns
			0xFF_00_00_00_00_00_00,    // Black pawns
			0x42,                      // White knights
			0x42_00_00_00_00_00_00_00, // Black knights
			0x24,                      // White bishops
			0x24_00_00_00_00_00_00_00, // Black bishops
			0x81,                      // White rooks
			0x81_00_00_00_00_00_00_00, // Black rooks
			0x08,                      // White queen
			0x08_00_00_00_00_00_00_00, // Black queen
			0x10,                      // White king
			0x10_00_00_00_00_00_00_00, // Black king
		},
		WhitesTurn:     true,
		CastlingRights: [4]bool{true, true, true, true},
		EnPassant:      0,
	}
}

// Move represents a chess move.
// It encodes the source square and destination square.
type Move struct{}

// NewMove creates a new move given a source and destination square.
func NewMove(source uint8, dest uint8) Move {
	panic("not implemented")
}

// LegalMoves defines a list of all legal moves for the position.
func (p Position) LegalMoves() []Move {
	panic("not implemented")
}

// IsLegal determines if a move is legal in the given position.
func (p Position) IsLegal(move Move) bool {
	panic("not implemented")
}
