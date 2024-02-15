package parser

import(
    "testing"
    "go-interpreter/ast"
    "go-interpreter/lexer"
)

func TestLetStatement(t *testing.T){
    input := `let x = 5;
    let y = 10;
    let foobar = 838383;`

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)
    if program == nil{
        t.Fatalf("program statements does not contain 3 statments. got=%d",
        len(program.Statements))
    }
    tests := []struct{
        expectedIdentifier string

    }{
        {"x"},
        {"y"},
        {"foobar"},
    }
    for i, tt := range tests{
        stmt := program.Statements[i]
        if !testLetStatement(t, stmt, tt.expectedIdentifier){
            return
        }
    }
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
    if s.TokenLiteral() != "let"{
        t.Errorf("not let.  got=%q", s.TokenLiteral())
    }
    letstmt, ok := s.(*ast.LetStatement)
    if !ok {
        t.Errorf("s not *ast.LetStatement. got=%T", s)
    }
    if letstmt.Name.Value != name{
        t.Errorf("letstmt.Name.Value not %s. got=%s", name, letstmt.Name.Value)
        return false
    }
    if letstmt.Name.TokenLiteral() != name {
        t.Errorf("letstmt.Name.TokenLiteral() not %s. got=%s", name, letstmt.Name.TokenLiteral())
        return false
    }
    return true
}

func checkParserErrors(t *testing.T, p *Parser){
    errors := p.Errors()
    if len(errors) == 0{
        return
    }
    t.Errorf("parser has %d errors", len(errors))
    for _, msg := range errors{
        t.Errorf("parser error: %q", msg)
    }
    t.FailNow()
}

func testReturnStatement(t *testing.T){
    input := `return 5;
    return 10;
    return 993322;`
    l := lexer.New(input)
    p := New(l)

    parse := p.ParseProgram()
    checkParserErrors(t, p)
    if len(parse.Statements) != 3{
        t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(parse.Statements))
    }
    for _, stmt := range parse.Statements {
        returnStatement, ok := stmt.(*ast.ReturnStatement)
        if !ok {
            t.Errorf("stmt not *ast.ReturnStatement. got=%T",stmt)
            continue
        }
        if returnStatement.TokenLiteral() != "return"{
            t.Errorf("returnStatement.TokenLiteral not 'return',got %q", returnStatement.TokenLiteral())
        }
    }
}
