# [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

```text
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
```


```text
(2 -> 4 -> 3) 342

(5 -> 6 -> 4) 465

(7 -> 0 -> 8) 807

342 + 465 = 807
```

Constraints:

* The number of nodes in each linked list is in the range ```[1, 100].```
* ```0 <= Node.val <= 9```
* It is guaranteed that the list represents a number that does not have leading zeros.
