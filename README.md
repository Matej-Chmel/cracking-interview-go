# Cracking the Coding Interview with Go

## Running a solution

To run a selected solution, change directory to its subfolder and run the `main.go` file. Example follows.

```cmd
cd arrays-and-strings/01-is-unique
go run ./main.go
```

Output will consist of a list of tested cases and comparison between actual result from the used solution function and expected result for that test case.

```none
withLoop(abc) -> true
OK!

withLoop(abcABC) -> true
OK!

withLoop(abc def) -> true
OK!

withLoop(abca) -> false
OK!

withLoop(AbcAb) -> false
OK!

withLoop(abc abc A) -> false
OK!

withArray(abc) -> true
OK!

withArray(abcABC) -> true
OK!

withArray(abc def) -> true
OK!

withArray(abca) -> false
OK!

withArray(AbcAb) -> false
OK!

withArray(abc abc A) -> false
OK!
```