package fundamentals

/*
Google interview
trellis
####
.#.#
.#..
.#.. <- ground level

oooo
.o.o
.X..
.#.. <- ground level

....
....
....
.#..

.###
##.#
.#.#
.X.#
.#..

.00#
.0.#
.0.#
.X.#
.#.#

. = empty trellis cell
# = trellis with live vine
X = cut point
o = 'doomed' vine segment
*/

// []string or [][]string

func PrintMyCut(trellis [][]string, cord []int) [][]string {

	// check inputs

	replace := "o"
	hasRoots := recurse(trellis, cord, replace)

	if !hasRoots {
		recurse(trellis, cord, ".")
	} else {
		recurse(trellis, cord, "#")
	}

	return trellis
}

func recurse(trellis [][]string, next []int, replace string) bool {

	rowlen := len(trellis)
	collen := len(trellis[0])
	if next[0] < 0 || next[1] < 0 {
		return false
	}

	if next[0] >= rowlen || next[1] >= collen {
		return false
	}

	if trellis[next[0]][next[1]] == "." || trellis[next[0]][next[1]] == replace {
		return false
	}

	if next[0] == (rowlen-1) && trellis[next[0]][next[1]] == "#" {
		return true
	}

	trellis[next[0]][next[1]] = replace

	//next
	return recurse(trellis, top, replace) ||
		recurse(trellis, left, replace) ||
		recurse(trellis, right, replace) ||
		recurse(trellis, down, replace)

}
