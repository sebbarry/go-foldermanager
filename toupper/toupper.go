package toupper

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"go-foldermanager/stack"
)


func ToUpperTraverse() error {
    fmt.Println("updating folders to upper case..")

    thisStack := &stack.Stack{}

    currPath, _ := os.Getwd()
    err := dfs(currPath, 1, thisStack) //get all the current paths.
    if err != nil {
        fmt.Println(err)
    }

    for !thisStack.IsEmpty() {
        node, _ := thisStack.Pop()
        fmt.Println(node.GetNodePath())

        err := changePaths(node)
        if err != nil {
            fmt.Println(err)
        }

    }
    return nil
}


func changePaths(n *stack.Node) error {
    oldPath := updatePath(n.GetNodePath(), n.GetNodeName())
    upDirName := updateName(n.GetNodeName())
    newPath := updatePath(n.GetNodePath(), upDirName)
    if oldPath == newPath {
        return errors.New("Same filenames")
    } else {
        fmt.Println("--------")
        fmt.Println("Changing:")
        fmt.Printf("old=%s", oldPath)
        fmt.Println()
        fmt.Printf("new=%s", newPath)
        err := os.Rename(oldPath, newPath)
        if err != nil {
            return err
        }
    }
    return nil

}

func dfs(currPath string, counter int, thisStack *stack.Stack) error {

    //get files from currentDir
    files, err := os.ReadDir(currPath)
    if err != nil {
        fmt.Println(err)
        return error(err)
    }

    if len(files) <= 0 {
        return errors.New("Reached end of path.")
    }

    //loop through all the files.
    for _, filename := range files {
        //update the currentpath name to include a folder.
        fullPath := updatePath(currPath, filename.Name())

        //Lstat -> get information from the current component..
        fi, err := os.Lstat(fullPath)
        if err != nil {
            fmt.Println(err)
            continue
        }


        //check if we have a directory.
        if fi.Mode().IsDir() {

            //add the folder to the thisStack.
            thisStack.Push(stack.FileType.DIR, fi.Name(), currPath)

            updatedCounter := counter+1
            dfs(fullPath, updatedCounter, thisStack)
        }

    }
    return nil

}


// ------------------------------------------
// ---------------- helpers -----------------
// ------------------------------------------




// **********
// input the whole path of the folder.
// **********
func updateName(folderName string) string {
    if len(folderName) <= 1 {
        //fmt.Println(strings.ToTitle(pathName))
    }
    upFolderName := strings.ToTitle(string(folderName[0])) + string(folderName[1:])
    return upFolderName
}

func updatePath(path string, newFolder string) string {
    return path + "/" + newFolder
}


func makeDelimiter(counter int) string {
    i := 0
    tmp := "-"
    final := ""
    for i <= counter {
        if(i+1 == counter) {
            final = final + "|" + tmp
        } else {
            final = final + tmp
        }
        i+=1
    }
    return final
}
