package stack

import (
	"fmt"
	"testing"
)


func TestPush(t *testing.T) {

    var expected = []string {
        FileType.DIR,
        FileType.FILE,
        FileType.DIR,
        FileType.FILE,
        FileType.DIR,
        FileType.FILE,
        FileType.DIR,
    }


    stack := &Stack{}


    for i:=0;i<=len(expected)-1;i++ {
        if i%2 == 0  {
            stack.Push(FileType.DIR, "tmp", "/test")
        } else {
            stack.Push(FileType.FILE, "tmp", "/test")
        }
    }

    j := 0
    for !stack.IsEmpty() {
        value, _ := stack.Pop()
        if expected[j] != value.nodeType {
            t.Errorf("expected value: %s. got: %s", value.nodeType, expected[j])
        }
        getInfo(value, stack)
        j+=1
    }
}


func getInfo(n *Node, s *Stack) {
    fmt.Printf("NodeType: %s, Stack Size: %d \n", n.nodeType, s.Size())
}

func TestPop(t *testing.T) {
    t.Log("succeeded.")
}
