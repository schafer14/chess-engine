package position_test

import (
	"testing"
	"unsafe"

	"github.com/schafer14/chess-engine/position"
)

// Success and failure markers.
const (
	success = "\u2713"
	failed  = "\u2717"
)

const (
	// testTypeExact specifies a test where the legal moves must exactly
	// match the result set provided by the test.
	testTypeExact int = iota
	// testTypeNone specifies a test where none of the moves in the result
	// set may be present in the list of generated moves.
	testTypeNone
	// testTypeRequired specifies a test where all moves in the result set
	// must be present in the generated set, but move may be as well.
	testTypeRequired
)

// testSpec defines a manual test that can be run.
type testSpec struct {
	should          string
	initialPosition string
	setupMoves      []string
	testType        int
	resultSet       []string
}

// tt defines a list of manual tests to run.
var tt = map[string][]testSpec{
	"When moving a king": {
		{
			should:          "should be able to move to any square adjacent to it",
			initialPosition: "8/8/8/3K4/8/8/8/k7 w KQkq - 0 1",
			testType:        testTypeExact,
			resultSet:       []string{"d4d5", "d4e5", "d4e4", "d4e3", "d4d3", "d4c3", "d4c4", "d4c5"},
		},
		{
			should:          "should not be able to move into its own piece",
			initialPosition: "8/8/8/3KR3/8/8/8/k7 w KQkq - 0 1",
			testType:        testTypeNone,
			resultSet:       []string{"d4e4"},
		},
		{
			should:          "should be able to capture an opponents piece",
			initialPosition: "8/8/8/3Kr3/8/8/8/k7 w KQkq - 0 1",
			testType:        testTypeRequired,
			resultSet:       []string{"d4e4"},
		},
		{
			should:          "should not be able to move into check",
			initialPosition: "8/8/8/3Krr2/8/8/8/k7 w KQkq - 0 1",
			testType:        testTypeNone,
			resultSet:       []string{"d4e4"},
		},
	},
}

func TestLegalMove(t *testing.T) {
	t.Log("Given a chess position")

	t.Log(unsafe.Sizeof(position.New()))
	t.Log(unsafe.Alignof(position.New()))

	for when, specs := range tt {
		t.Logf("\t%s", when)

		for _, spec := range specs {
			t.Logf("\t\tIt %s", spec.should)

			switch spec.testType {
			case testTypeExact:
				exactTest(t, spec)
			case testTypeNone:
				noneTest(t, spec)
			case testTypeRequired:
				requiredTest(t, spec)
			default:
				t.Fatalf("Unknown test type %d", spec.testType)
			}
		}
	}
}

func exactTest(t *testing.T, spec testSpec) {
	t.Errorf("\t\t\t%s: in position %q after %q\n\t\t\t\t\texpected %q\n\t\t\t\t\tgot %q", failed, spec.initialPosition, spec.setupMoves, spec.resultSet, []string{})
}

func noneTest(t *testing.T, spec testSpec) {
	t.Errorf("\t\t\t%s: in position %q after %q found unexpected move %q", failed, spec.initialPosition, spec.setupMoves, "e2e4")
}

func requiredTest(t *testing.T, spec testSpec) {
	t.Errorf("\t\t\t%s: in position %q after %q expected move %q not found", failed, spec.initialPosition, spec.setupMoves, "e2e4")
}
