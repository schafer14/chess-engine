package fen_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/schafer14/chess-engine/encoders/fen"
	"github.com/schafer14/chess-engine/position"
)

// Success and failure markers.
const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestUnmarshal(t *testing.T) {
	t.Log("Given the need to convert a fen string to a position.")
	t.Run("inital position", testInitialPosition)
	t.Run("arbitarty position", testArbitrary)
}

func testInitialPosition(t *testing.T) {
	t.Log("\tWhen the fen string is the initial position")

	fenString := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	expected := position.New()

	actual, err := fen.Unmarshal(fenString)
	if err != nil {
		t.Fatalf("\t\t%s It should not return an error: %s", failed, err)
	}
	t.Logf("\t\t%s It should not return an error", success)

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Fatalf("\t\t%s It should return the correct position\n\t\t\t%s", failed, diff)
	}
	t.Logf("\t\t%s It should return the correct position", success)
}

func testArbitrary(t *testing.T) {
	t.Log("\tWhen the fen string is an arbitary position")

	fenString := "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2"
	expected := position.New()
	expected.WhitesTurn = false
	expected.WhiteKnights = 0x20_00_02
	expected.WhitePawns = 0x10_00_EF_00
	expected.BlackPawns = 0xFB_00_04_00_00_00_00

	actual, err := fen.Unmarshal(fenString)
	if err != nil {
		t.Fatalf("\t\t%s It should not return an error: %s", failed, err)
	}
	t.Logf("\t\t%s It should not return an error", success)

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Fatalf("\t\t%s It should return the correct position\n\t\t\t%s", failed, diff)
	}
	t.Logf("\t\t%s It should return the correct position", success)
}
