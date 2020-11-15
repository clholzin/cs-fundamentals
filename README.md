# cs-fundamentals
computer science fundamentals

# Fundamentals:
* Binary Tree * recursion
* Array
* Stack
```
Stacker interface {

	IsEmpty() bool

	Peak() (int, error)

	Push(data int)

	Pop() (int, error)

}
```

* Queue
```
Queuer interface {

	IsEmpty() bool

	Peak() (int, error)

	Add(data int)

	Remove(data int) (int, error)

}
```
* Linked List
* Heaps
* Search Algo
* Hash Table
* Graphs

# Exercises:

## Rock Gravity

### Problem Description

Input:  A string which represents a two dimensional grid using the characters space, period, hyphen, colon, and newline.

Output:  The program should output a string which represents a two dimensional grid using the characters space, period, hyphen, colon, and newline.

The input characters have the following meaning:

- `.` period is a single rock
- `:` colon is two rocks
- `-` hyphen is a table
- ` ` space is empty space

The program should simulate gravity, pulling the rocks vertically down.  Rocks fall straight down until they hit the ground (the bottom of the grid) or a table (hyphen).  Rocks stack as densely as possible using the two-rock colon when there is more than one.

The output grid should be the same size as the input with the same number of lines.  The output string may omit trailing whitespace on each line.

### Examples

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

