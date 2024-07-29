# Cracking the Coding Interview with Go
Solutions in Go to questions in the book Cracking the Coding Interview.

## Running a solution
To run a selected solution, change directory to its subfolder and run the `main.go` file. Example follows.

```cmd
cd arrays-and-strings/1-is-unique
go run ./main.go
```

Output will consist of a list of tested cases and comparison between actual result from the used solution function and expected result for that test case.

```none
withArray
=========
(OK) abc -> true
(OK) abcABC -> true
(OK) abc def -> true
(OK) abca -> false
(OK) AbcAb -> false
(OK) abc abc A -> false

withLoop
========
(OK) abc -> true
(OK) abcABC -> true
(OK) abc def -> true
(OK) abca -> false
(OK) AbcAb -> false
(OK) abc abc A -> false
```