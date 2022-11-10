# Merge Sort

## The divide-and-conquer approach

**Recursive** algorithms typically follow a **divide-and-conquer** approach:
they break the problem into several subproblems that are similar to the original
problem but smaller in size, solve the subproblems recursively, and then combine
these solutions to create a solution to the original problem.

The divide-and-conquer paradigm involves three steps at each level of the recursion:

**Divide** the problem into a number of subproblems that are smaller instances
of the same problem.

**Conquer** the subproblems by solving them recursively. If the subproblems
sizes are small enough, just solve the subproblems in a straightforward manner.

**Combine** the solutions to the subproblems into the solution for the original
problem.

## Merge sort

The ***merge sort*** algorithm closely follows the divide-and-conquer paradigm.
Intuitively, it operates as follows:

**Divide:** Divide the *n*-element sequence to be sorted into two subsequences
of *n/2* elements each.

**Conquer:** Sort the two subsequences recursively using merge sort.

**Combine:** Merge the two sorted subsequences to produce the sorted answer.

The recursion "bottoms out" when the sequence to be sorted has length 1, in
which case there is no work to be done, since every sequence of length 1 is
already sorted.

![](mergeSort.gif)

## Quick Sort vs Merge Sort

1. **Partition of elements in the array**

   <details><summary>How are elements partitioned (ratio) in merge sort and quick sort?</summary>
   In the merge sort, the array is parted into just 2 halves (i.e. $n/2$).

   In case of quick sort, the array is parted into any ratio. There is no
   compulsion of dividing the array of elements into equal parts in quick sort.
   </details>

2. **Worst case complexity**

   <details><summary>What is the average and the worst case complexity of quick sort and merge sort?</summary>
   The worst case complexity of quick sort is $O(n^2)$ as there is need of lot of
   comparisons in the worst condition.

   In merge sort, worst case and average case has same complexities $O(n\log n)$.
   </details>

3. **Additional storage space requirement**

   <details><summary>Can quick sort and merge sort arrays/slices in-place?</summary>
   Merge sort is not in place because it requires additional memory space to
   store the auxiliary arrays.

   The quick sort is in place as it doesn’t require any additional storage.
   </details>

4. **Preferred for**

   <details><summary>Which sorting (quick sort/merge sort) is preferred for arrays/linked lists?</summary>
   Quick sort is preferred for arrays.

   Merge sort is preferred for linked lists.
   </details>

## Analyzing divide-and-conquer algorithms

When an algorithm contains a recursive call to itself, we can often describe its running time by a ***recurrence equation*** or ***recurrence***, which describes the overall running time on a problem of size *n* in terms of the running time on smaller inputs.

A recurrence for the running time of a divide-and-conquer algorithm falls out from the three steps of the basic paradigm. We let $T(n)$ be the running time on a problem of size $n$. If the problem size is small enough, say $n \leq c$ for some constant $c$, the straightforward solution takes constant time , which we write as $O(1)$. Suppose that our division of the problem yields $a$ subproblems, each of which is $1/b$ the size of the original. It takes time $T(n/b)$ to solve one subproblem of size $n/b$, and so it takes time $aT(n/b)$ to solve $a$ of them. If we take $D(n)$ time to divide the problem into subproblems and $C(n)$ time to combine the solutions to the subproblems into the solution to the original problem, we get recurrence

$$T(n) =
\begin{cases}
O(1) &\text{if }n \leq c,\\
aT(n/b) + D(n) + C(n) &\text{otherwise}.\\
\end{cases}$$

### Exercise 1

1. Calculate the values for $a$, $b$, $D(n)$ and $C(n)$ for the binary search algorithm and write the recurrence equation.

2. Solve the recurrence equation for the binary search algorithm. Explain why the
solution for the recurrence equation for the binary search is $O(\log n)$.

### Exercise 2

1. Calculate the values for $a$, $b$, $D(n)$ and $C(n)$ for the merge sort algorithm and write the recurrence equation.

2. Solve the recurrence equation for the merge sort. Explain why the solution for the recurrence equation for the merge sort is $O(n\log n)$.

### Exercise 3

We can express insertion sort as a recursive procedure as follows. In order to sort $A[1..n]$, we recursively sort $A[1..n-1]$ and then insert $A[n]$ into the sorted array $A[1..n-1]$.

1. Write a recurrence for the worst-case running time of this recursive version of insertion sort.
2. Solve the recurrence.
