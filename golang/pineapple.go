// https://github.com/karminski/pineapple

// backend.go
package pineapple 

import (
    "fmt"
    "errors"
)

type GlobalVariables struct {
    Variables  map[string]string
}

func NewGlobalVariables() *GlobalVariables {
    var g GlobalVariables
    g.Variables = make(map[string]string)
    return &g
}

func Execute(code string) {
    var ast *SourceCode
    var err  error 

    g := NewGlobalVariables()

    // parse
    if ast, err = parse(code); err != nil {
        panic(err)
    }

    // resolve
    if err = resolveAST(g, ast); err != nil {
        panic(err)
    }
}

func resolveAST(g *GlobalVariables, ast *SourceCode) error {
    if len(ast.Statements) == 0 {
        return errors.New("resolveAST(): no code to execute, please check your input.")
    }
    for _, statement := range ast.Statements {
        if err := resolveStatement(g, statement); err != nil {
            return err
        }
    }
    return nil
}

func resolveStatement(g *GlobalVariables, statement Statement) error {
    if assignment, ok := statement.(*Assignment); ok {
        return resolveAssignment(g, assignment)
    } else if print, ok := statement.(*Print); ok {
        return resolvePrint(g, print)
    } else {
        return errors.New("resolveStatement(): undefined statement type.")
    }
}

func resolveAssignment(g *GlobalVariables, assignment *Assignment) error {
    varName := "" 
    if varName = assignment.Variable.Name; varName == "" {
        return errors.New("resolveAssignment(): variable name can NOT be empty.")
    }
    g.Variables[varName] = assignment.String
    return nil
}

func resolvePrint(g *GlobalVariables, print *Print) error {
    varName := ""
    if varName = print.Variable.Name; varName == "" {
        return errors.New("resolvePrint(): variable name can NOT be empty.")
    }
    str := ""
    ok  := false
    if str, ok = g.Variables[varName]; !ok {
        return errors.New(fmt.Sprintf("resolvePrint(): variable '$%s'not found.", varName))
    }
    fmt.Print(str)
    return nil
}

// definition.go
package pineapple 

type Variable struct {
    LineNum int 
    Name    string 
}

type Assignment struct {
    LineNum   int 
    Variable *Variable
    String    string 
}

type Print struct {
    LineNum   int 
    Variable *Variable
}

type Statement interface{}

var _ Statement = (*Print)(nil)
var _ Statement = (*Assignment)(nil)

type SourceCode struct {
    LineNum      int 
    Statements []Statement
}

// lexer.go

package pineapple

import (
    "strings"
    "regexp"
    "fmt"
)

// token const
const (
    TOKEN_EOF         = iota  // end-of-file
    TOKEN_VAR_PREFIX          // $
    TOKEN_LEFT_PAREN          // (
    TOKEN_RIGHT_PAREN         // )
    TOKEN_EQUAL               // =
    TOKEN_QUOTE               // "
    TOKEN_DUOQUOTE            // ""
    TOKEN_NAME                // Name ::= [_A-Za-z][_0-9A-Za-z]*
    TOKEN_PRINT               // print
    TOKEN_IGNORED             // Ignored                
)

var tokenNameMap = map[int]string{
    TOKEN_EOF           : "EOF",
    TOKEN_VAR_PREFIX    : "$",
    TOKEN_LEFT_PAREN    : "(",
    TOKEN_RIGHT_PAREN   : ")",    
    TOKEN_EQUAL         : "=",
    TOKEN_QUOTE         : "\"",
    TOKEN_DUOQUOTE      : "\"\"",
    TOKEN_NAME          : "Name",
    TOKEN_PRINT         : "print",
    TOKEN_IGNORED       : "Ignored",
}

var keywords = map[string]int{
    "print"        : TOKEN_PRINT,
}

// regex match patterns
var regexName = regexp.MustCompile(`^[_\d\w]+`)

// lexer struct
type Lexer struct {
    sourceCode          string 
    lineNum             int    
    nextToken           string 
    nextTokenType       int 
    nextTokenLineNum    int
}

func NewLexer(sourceCode string) *Lexer {
    return &Lexer{sourceCode, 1, "", 0, 0} // start at line 1 in default.
}

func (lexer *Lexer) GetLineNum() int {
    return lexer.lineNum
}

func (lexer *Lexer) NextTokenIs(tokenType int) (lineNum int, token string) {
    nowLineNum, nowTokenType, nowToken := lexer.GetNextToken()
    // syntax error
    if tokenType != nowTokenType {
        err := fmt.Sprintf("NextTokenIs(): syntax error near '%s', expected token: {%s} but got {%s}.", tokenNameMap[nowTokenType], tokenNameMap[tokenType], tokenNameMap[nowTokenType]) 
        panic(err)
    }
    return nowLineNum, nowToken
}

func (lexer *Lexer) LookAheadAndSkip(expectedType int) {
    // get next token
    nowLineNum                := lexer.lineNum
    lineNum, tokenType, token := lexer.GetNextToken()
    // not is expected type, reverse cursor
    if tokenType != expectedType {
        lexer.lineNum              = nowLineNum
        lexer.nextTokenLineNum     = lineNum
        lexer.nextTokenType        = tokenType
        lexer.nextToken            = token
    }
}

func (lexer *Lexer) LookAhead() int {
    // lexer.nextToken* already setted
    if lexer.nextTokenLineNum > 0 {
        return lexer.nextTokenType
    }
    // set it
    nowLineNum                := lexer.lineNum
    lineNum, tokenType, token := lexer.GetNextToken()
    lexer.lineNum              = nowLineNum
    lexer.nextTokenLineNum     = lineNum
    lexer.nextTokenType        = tokenType
    lexer.nextToken            = token
    return tokenType
}

func (lexer *Lexer) nextSourceCodeIs(s string) bool {
    return strings.HasPrefix(lexer.sourceCode, s)
}

func (lexer *Lexer) skipSourceCode(n int) {
    lexer.sourceCode = lexer.sourceCode[n:]
}

func (lexer *Lexer) isIgnored() bool {
    isIgnored := false
    // target pattern
    isNewLine := func(c byte) bool {
        return c == '\r' || c == '\n'
    }
    isWhiteSpace := func(c byte) bool {
        switch c {
        case '\t', '\n', '\v', '\f', '\r', ' ':
            return true
        }
        return false
    }
    // matching
    for len(lexer.sourceCode) > 0 {
        if lexer.nextSourceCodeIs("\r\n") || lexer.nextSourceCodeIs("\n\r") {
            lexer.skipSourceCode(2)
            lexer.lineNum += 1
            isIgnored = true
        } else if isNewLine(lexer.sourceCode[0]) {
            lexer.skipSourceCode(1)
            lexer.lineNum += 1
            isIgnored = true
        } else if isWhiteSpace(lexer.sourceCode[0]) {
            lexer.skipSourceCode(1)
            isIgnored = true
        } else {
            break
        } 
    }
    return isIgnored
}

func (lexer *Lexer) scan(regexp *regexp.Regexp) string {
    if token := regexp.FindString(lexer.sourceCode); token != "" {
        lexer.skipSourceCode(len(token))
        return token
    }
    panic("unreachable!")
    return ""
}

// return content before token
func (lexer *Lexer) scanBeforeToken(token string) string {
    s := strings.Split(lexer.sourceCode, token)
    if len(s) < 2 {
        panic("unreachable!")
        return ""
    }
    lexer.skipSourceCode(len(s[0]))
    return s[0]
}

func (lexer *Lexer) scanName() string {
    return lexer.scan(regexName)
}

func (lexer *Lexer) GetNextToken() (lineNum int, tokenType int, token string) {
    // next token already loaded
    if lexer.nextTokenLineNum > 0 {
        lineNum                = lexer.nextTokenLineNum
        tokenType              = lexer.nextTokenType
        token                  = lexer.nextToken
        lexer.lineNum          = lexer.nextTokenLineNum
        lexer.nextTokenLineNum = 0
        return
    }
    return lexer.MatchToken()

}

func (lexer *Lexer) MatchToken() (lineNum int, tokenType int, token string) {
    // check ignored
    if lexer.isIgnored() {
        return lexer.lineNum, TOKEN_IGNORED, "Ignored"
    }
    // finish
    if len(lexer.sourceCode) == 0 {
        return lexer.lineNum, TOKEN_EOF, tokenNameMap[TOKEN_EOF]
    }
    // check token
    switch lexer.sourceCode[0] {
    case '$' :
        lexer.skipSourceCode(1)
        return lexer.lineNum, TOKEN_VAR_PREFIX, "$"
    case '(' :
        lexer.skipSourceCode(1)
        return lexer.lineNum, TOKEN_LEFT_PAREN, "("
    case ')' :
        lexer.skipSourceCode(1)
        return lexer.lineNum, TOKEN_RIGHT_PAREN, ")"
    case '=' :
        lexer.skipSourceCode(1)
        return lexer.lineNum, TOKEN_EQUAL, "="
    case '"' :
        if lexer.nextSourceCodeIs("\"\"") {
            lexer.skipSourceCode(2)
            return lexer.lineNum, TOKEN_DUOQUOTE, "\"\""
        }
        lexer.skipSourceCode(1)
        return lexer.lineNum, TOKEN_QUOTE, "\""
    }
    // check multiple character token
    if lexer.sourceCode[0] == '_' || isLetter(lexer.sourceCode[0]) {
        token := lexer.scanName()
        if tokenType, isMatch := keywords[token]; isMatch {
            return lexer.lineNum, tokenType, token
        } else {
            return lexer.lineNum, TOKEN_NAME, token
        }
    }
    // unexpected symbol
    err := fmt.Sprintf("MatchToken(): unexpected symbol near '%q'.", lexer.sourceCode[0])
    panic(err)
    return 
}

func isLetter(c byte) bool {
    return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

// parser.go

package pineapple

import (
    "errors"
)

// Name ::= [_A-Za-z][_0-9A-Za-z]*
func parseName(lexer *Lexer) (string, error) {
    _, name := lexer.NextTokenIs(TOKEN_NAME)
    return name, nil
}

// String ::= '"' '"' Ignored | '"' StringCharacter '"' Ignored
func parseString(lexer *Lexer) (string, error) {
    str := "" 
    switch lexer.LookAhead() {
    case TOKEN_DUOQUOTE:
        lexer.NextTokenIs(TOKEN_DUOQUOTE)
        lexer.LookAheadAndSkip(TOKEN_IGNORED)
        return str, nil 
    case TOKEN_QUOTE:
        lexer.NextTokenIs(TOKEN_QUOTE)
        str = lexer.scanBeforeToken(tokenNameMap[TOKEN_QUOTE])
        lexer.NextTokenIs(TOKEN_QUOTE)
        lexer.LookAheadAndSkip(TOKEN_IGNORED)
        return str, nil
    default:
        return "", errors.New("parseString(): not a string.")
    }
} 

// Variable ::= "$" Name Ignored
func parseVariable(lexer *Lexer) (*Variable, error) {
    var variable Variable
    var err      error 
    
    variable.LineNum = lexer.GetLineNum()
    lexer.NextTokenIs(TOKEN_VAR_PREFIX)
    if variable.Name, err = parseName(lexer); err != nil {
        return nil, err
    }
    lexer.LookAheadAndSkip(TOKEN_IGNORED)
    return &variable, nil
}

// Assignment  ::= Variable Ignored "=" Ignored String Ignored
func parseAssignment(lexer *Lexer) (*Assignment, error) {
    var assignment Assignment
    var err        error
    
    assignment.LineNum = lexer.GetLineNum()
    if assignment.Variable, err = parseVariable(lexer); err != nil {
        return nil, err
    }
    lexer.LookAheadAndSkip(TOKEN_IGNORED)
    lexer.NextTokenIs(TOKEN_EQUAL)
    lexer.LookAheadAndSkip(TOKEN_IGNORED)
    if assignment.String, err = parseString(lexer); err != nil {
        return nil, err
    }
    lexer.LookAheadAndSkip(TOKEN_IGNORED)
    return &assignment, nil
}

// Print ::= "print" "(" Ignored Variable Ignored ")" Ignored
func parsePrint(lexer *Lexer) (*Print, error) {
    var print Print 
    var err   error

    print.LineNum = lexer.GetLineNum()
    lexer.NextTokenIs(TOKEN_PRINT)
    lexer.NextTokenIs(TOKEN_LEFT_PAREN)
    lexer.LookAheadAndSkip(TOKEN_IGNORED)
    if print.Variable, err = parseVariable(lexer); err != nil {
        return nil, err
    }
    lexer.LookAheadAndSkip(TOKEN_IGNORED)
    lexer.NextTokenIs(TOKEN_RIGHT_PAREN)
    lexer.LookAheadAndSkip(TOKEN_IGNORED)
    return &print, nil
}


// Statement ::= Print | Assignment
func parseStatements(lexer *Lexer) ([]Statement, error) {
    var statements []Statement 
    
    for !isSourceCodeEnd(lexer.LookAhead()) {
        var statement Statement
        var err       error
        if statement, err = parseStatement(lexer); err != nil {
            return nil, err 
        }
        statements = append(statements, statement)
    }
    return statements, nil
}

func parseStatement(lexer *Lexer) (Statement, error) {
    lexer.LookAheadAndSkip(TOKEN_IGNORED) // skip if source code start with ignored token
    switch lexer.LookAhead() {
    case TOKEN_PRINT:
        return parsePrint(lexer)
    case TOKEN_VAR_PREFIX:
        return parseAssignment(lexer)
    default:
        return nil, errors.New("parseStatement(): unknown Statement.")
    }
}

// SourceCode ::= Statement+ 
func parseSourceCode(lexer *Lexer) (*SourceCode, error) {
    var sourceCode SourceCode
    var err        error

    sourceCode.LineNum = lexer.GetLineNum()
    if sourceCode.Statements, err = parseStatements(lexer); err != nil {
        return nil, err
    }
    return &sourceCode, nil
}

func isSourceCodeEnd(token int) bool {
    if token == TOKEN_EOF {
        return true
    }
    return false
}

func parse(code string) (*SourceCode, error) {
    var sourceCode *SourceCode
    var err         error 

    lexer := NewLexer(code)
    if sourceCode, err = parseSourceCode(lexer); err != nil {
        return nil, err 
    }
    lexer.NextTokenIs(TOKEN_EOF)
    return sourceCode, nil
}