package codegen

import (
    "fmt"
    "strings"
)

type CodeLines struct {
    lines []string
}

func NewCodeLines() *CodeLines {
    return &CodeLines{lines: make([]string, 0)}
}

func (c *CodeLines) Push(s string, a...interface{}) {
    c.lines = append(c.lines, fmt.Sprintf(s, a...))
}

func (c *CodeLines) ToString() string {
    return strings.Join(c.lines, "\n")
}

func main() {
}
