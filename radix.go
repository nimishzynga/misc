package main

import "strings"
import "fmt"

type TreeNode struct {
    prefix string
    isKey  bool
    childs []*TreeNode
}

func NewTree() *TreeNode {
    return CreateNodeWithSetKey("/")
}

func CreateNode(str string) *TreeNode {
    fmt.Println("allocating new node for ",str)
    return &TreeNode{
        prefix: str,
        childs : make([]*TreeNode, 0),
    }
}

func (t *TreeNode) SetKey() {
    t.isKey = true
}

func CreateNodeWithSetKey(str string) *TreeNode {
    t := CreateNode(str)
    t.SetKey()
    return t
}

func common(a,b string) int {
    fmt.Println("in common", a, b)
    var i int
    for i = range a {
        if a[i] != b[i] {
            fmt.Println("matching", i, b[i])
            break
        }
    }
    fmt.Println("i is ", i)
    return i
}

func (t *TreeNode) AddKey(key string) error {
    fmt.Println("addkey with len", key)
    if t == nil {
       return fmt.Errorf("Root is nil")
    }
    for _,ch := range t.childs {
        if strings.HasPrefix(key, ch.prefix) {
            return ch.AddKey(key[len(ch.prefix):])
        } else if val := common(key, ch.prefix); val > 0 {
            fmt.Println("common is not zero", val)
            nn := CreateNode(ch.prefix[val:])
            if ch.isKey {
                nn.SetKey()
                ch.isKey = false
            }
            nn.childs,ch.childs = ch.childs,nn.childs
            ch.childs = append(ch.childs, nn)
            ch.prefix = ch.prefix[:val]
            return ch.AddKey(key[val:])
        }
    }
    t.childs = append(t.childs, CreateNodeWithSetKey(key))
    return nil
}

func (t *TreeNode) DelKey(key string) error {
    if key == "" {
        return nil
    }
    if t == nil {
       return fmt.Errorf("Root is nil")
    }
    for _,ch := range t.childs {
        if strings.HasPrefix(key,ch.prefix) {
            ch.DelKey(key[len(ch.prefix):])
        }
    }
    if len(t.childs) == 1 && !t.isKey {
        t.prefix = t.prefix + t.childs[0].prefix
        t.childs = t.childs[0].childs
    }
    return nil
}

func (t *TreeNode) GetKeyPrefix(p string, full bool) ([]string) {
    fmt.Println("entering GetKeyPrefix", p)
    var ret,keys []string
    if t == nil {
        return ret
    }
    for k,ch := range t.childs {
        fmt.Println(ch.prefix, p, k)
        keys = nil
        found := true
        if full {
            keys = ch.GetKeyPrefix(p, full)
        }  else if strings.HasPrefix(ch.prefix, p) {
            fmt.Println("inside 2")
            keys = ch.GetKeyPrefix(p, true)
        } else if strings.HasPrefix(p, ch.prefix) {
            fmt.Println("inside 1")
            keys = ch.GetKeyPrefix(p[len(ch.prefix):], full)
        } else {
            found = false
        }
        if found && len(ch.childs) == 0 {
                ret = append(ret, ch.prefix)
        } else {
            for _,k := range keys {
                ret = append(ret, ch.prefix + k)
            }
            if len(keys) > 0 && ch.isKey && len(ch.prefix) >= len(p) {
                ret = append(ret, ch.prefix)
            }
        }
    }
    fmt.Println("extinging GetKeyPrefix", p)
    return ret
}

func main() {
    nt := NewTree()
    nt.AddKey("mit")
    nt.AddKey("mib")
    nt.AddKey("mig")
    fmt.Println(nt.GetKeyPrefix("mi", false))
}
