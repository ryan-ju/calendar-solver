# calendar-solver
Solves the calendar puzzle

# Run

To solve 24th May, run
```
go run ./cmd solve 0524
```

# Run In Web

Copy [single_file.go](single_file.go) into https://goplay.space/ and run.

Or visit this https://play.golang.org/p/sbIpoGEVIxn

It prints out like this:
```
searched times: 25876, duration = 0ms, #solutions = 59
solutions = 
| S | S | C | C | C | Z | . |
| T | S | C | . | C | Z | . |
| T | S | S | V | L | Z | Z |
| T | T | . | V | L | P | Z |
| T | V | V | V | L | P | P |
| O | O | O | L | L | P | P |
| O | O | O | . | . | . | . |

# more solutions ...
```

# Also See

https://github.com/zjuasmn/calendar-puzzle-solver uses the same DFS (depth-first-search) algorithm.