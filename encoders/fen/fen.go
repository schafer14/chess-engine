package fen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/schafer14/chess/position"
)

var (
	// ErrInvalidFEN is used when an inproper fen value is passed to the unmarshaller.
	ErrInvalidFEN = fmt.Errorf("invalid fen")
)

// Unmarshal turns a fen string into a postion that can be used for computation.
func Unmarshal(fenString string) (position.Position, error) {
	var p position.Position

	// General check that the fen string has enough parts.
	// This prevents panics in the case that an invalid FEN string is provided
	parts := strings.Split(fenString, " ")
	if len(parts) < 4 {
		return p, ErrInvalidFEN
	}

	// Parse the turn segment of the FEN string
	if parts[1] == "w" {
		p.WhitesTurn = true
	}

	// Parse the castling privledges segment of the fen string
	if strings.Contains(parts[2], "K") {
		p.CastlingRights[0] = true
	}

	if strings.Contains(parts[2], "Q") {
		p.CastlingRights[1] = true
	}

	if strings.Contains(parts[2], "k") {
		p.CastlingRights[2] = true
	}

	if strings.Contains(parts[2], "q") {
		p.CastlingRights[3] = true
	}

	// Parse the en passant segement of the FEN string
	if parts[3] == "-" {
		p.EnPassant = 0
	} else {
		enPassant64, err := strconv.ParseInt(parts[3], 10, 8)
		if err != nil {
			return p, ErrInvalidFEN
		}

		p.EnPassant = uint8(enPassant64)
	}

	// Parse the position segment of the FEN string.
	positionParts := strings.Split(parts[0], "/")
	if len(positionParts) != 8 {
		return p, ErrInvalidFEN
	}

	for i := 0; i < 8; i++ {
		rowIndex := 0
		for _, c := range positionParts[7-i] {
			sqNum := i*8 + rowIndex

			switch c {
			case 'P':
				p.WhitePawns |= 1 << sqNum
				rowIndex += 1
			case 'N':
				p.WhiteKnights |= 1 << sqNum
				rowIndex += 1
			case 'B':
				p.WhiteBishops |= 1 << sqNum
				rowIndex += 1
			case 'R':
				p.WhiteRooks |= 1 << sqNum
				rowIndex += 1
			case 'Q':
				p.WhiteQueens |= 1 << sqNum
				rowIndex += 1
			case 'K':
				p.WhiteKing = uint8(sqNum + 1)
				rowIndex += 1
			case 'p':
				p.BlackPawns |= 1 << sqNum
				rowIndex += 1
			case 'n':
				p.BlackKnights |= 1 << sqNum
				rowIndex += 1
			case 'b':
				p.BlackBishops |= 1 << sqNum
				rowIndex += 1
			case 'r':
				p.BlackRooks |= 1 << sqNum
				rowIndex += 1
			case 'q':
				p.BlackQueens |= 1 << sqNum
				rowIndex += 1
			case 'k':
				p.BlackKing = uint8(sqNum + 1)
				rowIndex += 1
			default:
				skip, err := strconv.ParseInt(string(c), 10, 64)
				if err != nil {
					return p, ErrInvalidFEN
				}
				rowIndex += int(skip)
			}

		}
	}

	return p, nil
}
