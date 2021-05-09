package position

// Position represents a chess position.
// A position can be used to generate a list of legal moves or validate a move is legal.
type Position struct {
	WhitePawns   uint64
	BlackPawns   uint64
	WhiteKnights uint64
	BlackKnights uint64
	WhiteBishops uint64
	BlackBishops uint64
	WhiteRooks   uint64
	BlackRooks   uint64
	WhiteQueens  uint64
	BlackQueens  uint64

	// We only allow for a maximum of one king per side on the board at once.
	// This increases board copy speeds by about 20%.
	WhiteKing uint8
	BlackKing uint8

	WhitesTurn     bool
	CastlingRights [4]bool
	EnPassant      uint8
}

// New creates a new starting position.
func New() Position {
	return Position{

		WhitePawns:   0xFF_00,                   // White pawns
		BlackPawns:   0xFF_00_00_00_00_00_00,    // Black pawns
		WhiteKnights: 0x42,                      // White knights
		BlackKnights: 0x42_00_00_00_00_00_00_00, // Black knights
		WhiteBishops: 0x24,                      // White bishops
		BlackBishops: 0x24_00_00_00_00_00_00_00, // Black bishops
		WhiteRooks:   0x81,                      // White rooks
		BlackRooks:   0x81_00_00_00_00_00_00_00, // Black rooks
		WhiteQueens:  0x08,                      // White queen
		BlackQueens:  0x08_00_00_00_00_00_00_00, // Black queen

		WhiteKing: 5,  // White king
		BlackKing: 61, // Black king

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
func LegalMoves(p Position) []Move {
	panic("not implemented")
}

// IsLegal determines if a move is legal in the given position.
func IsLegal(p Position, move Move) bool {
	panic("not implemented")
}
