package lists

import (
	"fmt"
	"testing"
)

func TestListX(t *testing.T) {
	ll := NewListX[string, string]()
	fmt.Println("------------------AppendHead")
	ll.AppendHead(&ValueStu[string, string]{Key: "9999", Value: "gogo"})
	e := ll.GetTailValue()
	fmt.Println("GetTailValue-----", *e)
	fmt.Println("------------------AppendTail")
	ll.AppendTail(&ValueStu[string, string]{Key: "8888", Value: "gogo1"})
	e = ll.GetTailValue()
	fmt.Println("GetTailValue-----", *e)
	fmt.Println("------------------PopTail")
	ll.PopTail()
	e = ll.GetTailValue()
	fmt.Println("GetTailValue-----", *e)
	fmt.Println("------------------PopTail")
	ll.PopTail()
	e = ll.GetTailValue()
	if e != nil {
		fmt.Println(*e)
	}
	fmt.Println("------------------GetValuesByKey")
	ll.AppendTail(&ValueStu[string, string]{Key: "8888", Value: "gogo2"})
	ll.AppendTail(&ValueStu[string, string]{Key: "8888", Value: "gogo3"})
	ll.AppendTail(&ValueStu[string, string]{Key: "8888", Value: "gogo4"})
	ll.AppendTail(&ValueStu[string, string]{Key: "8888", Value: "gogo5"})
	v2 := &ValueStu[string, string]{Key: "8888", Value: "gogo3"}
	fmt.Println("------------------AppendTail")
	ll.AppendTail(v2)
	fmt.Println("------------------PopByKey")
	ll.PopByKey(v2)
	es := ll.GetValuesByKey("8888")
	for i := range es {
		fmt.Println(*es[i])
	}
}

func TestListX_AppendHead(t *testing.T) {
	ll := NewListX[string, string]()
	fmt.Println("------------------AppendHead")
	ll.AppendHead(&ValueStu[string, string]{Key: "9999", Value: "gogo"})
	e := ll.GetTailValue()
	fmt.Println("GetTailValue-----", *e)
}

func TestListX_AppendTail(t *testing.T) {
	ll := NewListX[string, string]()
	fmt.Println("------------------AppendHead")
	ll.AppendTail(&ValueStu[string, string]{Key: "9999", Value: "gogo"})
	e := ll.GetTailValue()
	fmt.Println("GetTailValue-----", *e)
}

func TestListX_AppendBefore(t *testing.T) {
	ll := NewListX[string, string]()
	fmt.Println("------------------AppendHead")
	ll.AppendTail(&ValueStu[string, string]{Key: "9999", Value: "gogo"})
	e := ll.GetTailElement()
	fmt.Println("GetTailValue-----", *e)
	v2 := &ValueStu[string, string]{Key: "8989", Value: "dfafaf"}
	ll.AppendBefore(v2, e)
	fmt.Println(*ll.GetHeadValue())
}
