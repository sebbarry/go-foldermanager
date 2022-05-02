package stack

import (
	"errors"
)


type Node struct {
	nodeType string
	nodeName string
    nodePath     string
	next     *Node
}
func (n *Node) GetNodeName() string {
    return n.nodeName
}
func (n *Node) GetNodePath() string {
    return n.nodePath
}
func (n *Node) GetNodeType() string {
    return n.nodeType
}
func (n *Node) PeekNext() *Node {
    return n.next
}


type fileType = string
var FileType = struct {
    DIR      fileType
    FILE     fileType
    SYMLINK  fileType
}{
    DIR: "directory",
    FILE: "file",
    SYMLINK: "symboliclink",
}

type Stack struct {
	top     *Node
	isEmpty bool
	size    int
}

func (s *Stack) Size() int { return s.size }
func (s *Stack) IsEmpty() bool { return s.isEmpty }




func (s *Stack) Push(nodeType fileType, nodeName string, nodePath string) {
    newNode :=&Node{nodeType: nodeType, nodeName: nodeName, nodePath: nodePath}
    if s.top == nil {
        s.top = newNode
    } else {
        tmp := s.top
        head := newNode
        head.next = tmp
        s.top = head
    }
    s.size += 1
    s.isEmpty = false
}



func (s *Stack) Peek() (*Node, error) {
	if s.Size() <= 0 {
		return nil, errors.New("Empty stack.")
	} else {
		tmp := s.top
		return tmp, nil
	}
}




func (s *Stack) Pop() (*Node, error) {
	if s.Size() <= 0 {
        s.isEmpty = true
		return nil, errors.New("Empty stack")
	}
    if s.Size() == 1 {
        s.size = 0
        s.isEmpty = true
        tmp := s.top
        s.top = nil
        return tmp, nil
    }

    s.isEmpty = false
    s.size -= 1

    tmp := s.top
	s.top = s.top.next

	return tmp, nil
}
