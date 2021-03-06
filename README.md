# Insertion Sort Algorithm

<img alt="from js to go image" src="https://github.com/algorithms-go/insertion-sort/blob/master/thumb.jpg?raw=true" width="600"/>

### Overview

Insertion sorting is one of the simplest sorting algorithms, which
also proves to be one of the fastest for small sized data sets, having
an average time complexity of: O(k*n).

Insertion Sorting works pretty much like sorting a hand of cards.
The idea is pretty much the same, we take each card and
insert it into its correct position, till the whole hand is
sorted.

Here are the basic steps of the algorithm:

- Iterate from **LEFT** to **RIGHT**
- **INSERT** each element into its **CORRECT POSITION**
- The **LEFT** side (before the key) is always **SORTED**

### The Algorithm

Let's suppose we have small array/slice of 6 numbers:

[`5`,`2`,`4`,`6`,`1`,`3`]

Here's how the ***Insertion Sorting Algorithm*** will work step by step:

<img alt="from js to go image" src="https://github.com/algorithms-go/insertion-sort/blob/master/insertion-sort.gif?raw=true" width="600"/>


Sorting it using Insertion Sort will result in the
following steps:

**0)** `5` ***2*** `4` `6` `1` `3`

**1)** `2` `5` ***4*** `6` `1` `3`

**2)** `2` `4` `5` ***6*** `1` `3`

**3)** `2` `4` `5` `6` ***1*** `3`

**4)** `1` `2` `4` `5` `6` ***3***

**5)** `1` `2` `3` `4` `5` `6`

### Pseudocode

```
for j = 1 to A.length
  key = A[j]
  i = j-1
  while i>-1 and A[i] > key
    A[i+i] = A[i]
    i = i-1
  A[i+1] = key
```

### Facts

- **Simple** to implement
- **Efficient** and **Fastest** for small data sets
- Time Complexity: O(k*n)
- Best: O(n)
- Worst: O(n*n)

### Resources

- [Go Standard Library - sort package](https://github.com/golang/go/blob/master/src/sort/sort.go#L196)
- [Medium Article - Insertion Sorting in Go](https://steevehook.medium.com/insertion-sorting-algorithm-in-go-26f1ec49936c)
- [YouTube Video - Insertion Sorting in Go](https://youtu.be/vg4GS4LSJXI)
