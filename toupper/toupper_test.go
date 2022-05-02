package toupper

import (
    "testing"
    "fmt"
)



func TestUpper(t *testing.T) {
    val := ToUpperTraverse()
    fmt.Println(val)
    if val != nil {
        t.Error(val)
    }
    return
}



//function to test for errors in our parser.
//func checkParserErrors(t *testing.T, p *Parser)  {
//    errors := p.Errors()
//    if len(errors) == 0 {
//        return
//    }
//    t.Errorf("Parser has %d errors", len(errors))
//    for _, msg := range errors {
//        t.Errorf("parser error: %q", msg)
//    }
//    t.FailNow()
//}


