# cs-fundamentals
computer science fundamentals

# Fundamentals:
* Binary Tree  
  ```go
  Treer interface {
  	Search(int) *Tree
  	Insert(int, *Tree) *Tree
  	Remove(int) bool
  	Traverse()
  }
  ```
* Array
  ```go
  Lister interface {
  	Push(Value) bool
  	Pop() bool
  	Shift() bool
  	FindIndex(Value) (Value, bool)
  	Value(index int) (Value, bool)
  	Count() int
  	DeleteAt(index int) (Value, bool)
  	Iterate() []Value
  }
  ```
* Linked List

  * Stack
    ```go
    Stacker interface {
    	IsEmpty() bool
    	Peak() (int, error)
    	Push(data int)
    	Pop() (int, error)
    }
    ```

  * Queue
    ```go
    Queuer interface {
    	IsEmpty() bool
    	Peak() (int, error)
    	Add(data int)
    	Remove(data int) (int, error)
    }
    ```
* Search Algo
  * Binary Search
  * Depth-First Search
  * Bredth-First Search
* Hash Min Table
   ```
   type MinIntHeap struct {
   	 Cap   int
   	 Size  int
   	 items []int
    }
   ```
* Heaps 
* Graphs
* Sorting
    * MergeSort(arr[], l,  r)
      ```If r > l
         1. Find the middle point to divide the array into two halves:  
                 middle m = (l+r)/2
         2. Call mergeSort for first half:   
                 Call mergeSort(arr, l, m)
         3. Call mergeSort for second half:
                 Call mergeSort(arr, m+1, r)
         4. Merge the two halves sorted in step 2 and 3:
                 Call merge(arr, l, m, r)
      ```
    * heapSort
    * quickSort
    * Distribution Sort -> selectionSort

# Exercises:

## Subtract

Given a singly linked list, modify the value of first half nodes such that :

1st node’s new value = the last node’s value - first node’s current value
2nd node’s new value = the second last node’s value - 2nd node’s current value,
and so on …

 NOTE :
If the length L of linked list is odd, then the first half implies at first floor(L/2) nodes. So, if L = 5, the first half refers to first 2 nodes.
If the length L of linked list is even, then the first half implies at first L/2 nodes. So, if L = 4, the first half refers to first 2 nodes.
Example :

Given linked list 1 -> 2 -> 3 -> 4 -> 5,

You should return 4 -> 2 -> 3 -> 4 -> 5
as

for first node, 5 - 1 = 4
for second node, 4 - 2 = 2
Try to solve the problem using constant extra space.




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

