package fundamentals

func roll(rows [][]string, rowIndex, columnIndex int) {

	above := rowIndex - 1
	if above < 0 {
		return
	}
	current := rows[rowIndex][columnIndex]
	aboveValue := rows[above][columnIndex]

	if somethingToMove(aboveValue) {
		if goodToMove(current) {
			rows[rowIndex][columnIndex] = aboveValue
			rows[rowIndex][above] = " "
			if rowIndex != len(rows) {
				rowIndex-- // roll it down
				roll(rows, rowIndex, columnIndex)
			} else {
				roll(rows, above, columnIndex)
			}
		}
	} else if above != 0 {
		roll(rows, above, columnIndex)
	}

}

func goodToMove(value string) bool {

	switch value {
	case " ":
	case ".":
		return true
	case ":":
		return false
	default:
		break
	}
	return false
}

func somethingToMove(value string) bool {

	switch value {
	case " ":
	case "-":
		return false
	case ".":
	case ":":
		return true
	default:
		break
	}
	return false
}

/*

  for i :=len(rows)-1; i>0; i--{





    nextIndex = index + 1
    if nextIndex > len(rows){
      break
    }
    next := rows[nextIndex]

    columns := strings.Split(current,"")
    nextColumns := strings.Split(next,"")

    for k,v := range columns {


      if decide(v) {
        if decide(nextColumns[k])
      }


      /*
      current positoiin
      next position
         if table, period->colon, space->replace




    }
  }    */

// fmt.Println(data)*/
//}

/*
Your previous Markdown content is preserved below:

## Problem Description

Input:  A string which represents a two dimensional grid using the characters space, period, hyphen, colon, and newline.

Output:  The program should output a string which represents a two dimensional grid using the characters space, period, hyphen, colon, and newline.

The input characters have the following meaning:

- `.` period is a single rock
- `:` colon is two rocks
- `-` hyphen is a table
- ` ` space is empty space

The program should simulate gravity, pulling the rocks vertically down.  Rocks fall straight down until they hit the ground (the bottom of the grid) or a table (hyphen).  Rocks stack as densely as possible using the two-rock colon when there is more than one.

The output grid should be the same size as the input with the same number of lines.  The output string may omit trailing whitespace on each line.

## Examples

1. A single rock falls down:

Input:
```
.


```
Output:
```


.
```

2. Two single rocks fall into a dense two rock stack:

Input:
```
.
.
```
Output:
```

:
```

3. Five rocks fall into a stack of two colons and a single period:

Input:
```
.
:

.
.
```
Output:
```


.
:
:
```

4. Rock cannot fall through a table:

Input:
```
.
.
-
.

```
Output:
```

:
-

.
```

5. Each column is an independent stack of rocks:

Input:
```
.      .
: .  :
-   .  -
. -
.    .
```
Output:
```
.
:      .
- .    -
  -  .
:   .:
```

*/
