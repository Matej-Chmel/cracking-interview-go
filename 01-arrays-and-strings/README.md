# Arrays and strings

## Is unique?
Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?

|Time|Space|Algorithm|
|:--|:--|:--|
| O(n) | O(1) | Use bool array or bitset. Convert each character to an index and check if its position is already set. |
| O(n²) | O(1) | Use two nested loops. For each character check if all characters to the right from it are different. |

## Check permutation
Given two strings, write a method to decide if one is a permutation of the other.

|Time|Space|Algorithm|
|:--|:--|:--|
| O(n·log&nbsp;n) | O(1) | Sort each string and check if they are equal. |
| O(n) | O(1) | Use integer array and 3 loops. First, iterate over first string and increase position of current character in the array by one. Second, iterate over second string and decrease position of current character by one. Third, strings are permutations if all positions are zero. |

## URLify
Write a method to replace all spaces in a string with *%20*. The string has sufficient amount of spaces at the end.

|Time|Space|Algorithm|
|:--|:--|:--|
| O(n) | O(1) | Use two pointers. Set the left pointer to the first non-space character and the right to the last character. Iterate backwards. Write characters on a position of the left pointer to the position of the right pointer. Spaces are converted to three characters *%20*. |

## Palindrome permutation
Given a string, write a function to check if it is a permutation of a palindrome.

|Time|Space|Algorithm|
|:--|:--|:--|
| O(n) | O(1) | Use bool array and two loops. Iterate over the string and flip the switch at position of the current character. Then check that only one or zero switches are flipped which means that only one or zero characters are present odd number of times. |
| O(n) | O(1) | Use bitset and one loop. The algorithm for the first loop is the same. The second loop is replaced with operation `(bits & (bits - 1)) == 0` which checks whether there is any overlap between the original bitset and bitset after subtracting a one. Bitset with only one bit set will have no overlap. |

## One away
Given two strings, write a function to check if they are one or zero edits away. An edit is an insert, removal or replacement of a character.

|Time|Space|Algorithm|
|:--|:--|:--|
| O(n) | O(1) | Use two pointers. First, set shorter and longer string and their pointers to zero index. Then iterate over pairs of characters. If the characters are the same, move both pointers. If they differ, move long pointer but the short one only if the lengths of both strings match. If two pairs differ, more than one edit is required so return false. When all characters are processed, return true. |

## String compression
Implement a method to perform basic string compression using the counts of repeated characters. If the compressed string would not become smaller than the original string, return the original.

|Time|Space|Algorithm|
|:--|:--|:--|
| O(n²) | O(n) | Use a string builder, one character and one integer. Iterate over characters and count the number of occurrences of the previous character. When character changes, write it and its amount to the string builder. |
| O(n²) | O(n) | First count the length of the compressed string using comparison versus the next power of 10. If the length is shorter than the original, grow the string builder to that length and execute the previous algorithm. |
| O(n²) | O(n) | Use a string builder and an integer array. First write characters as integers and their counts into the array. At the same time compute the length of the compressed string. Then compare the length against the original. If the compressed string is shorter, grow the string builder and write the information from the array to the builder. |