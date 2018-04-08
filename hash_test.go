package consistenthash

import (
	"strings"
	"testing"
)

func TestRing(t *testing.T) {
	ring := NewRing(50)
	ring.AddNodes([]string{"node1", "node2", "node3", "node3"})

	if ring.Nodes.Len() != 200 {
		t.Error("it shoud have 200 nodes")
	}
}

func TestRemoveNode(t *testing.T) {
	ring := NewRing(50)
	ring.AddNodes([]string{"node1", "node2", "node3", "node3"})

	err := ring.RemoveNodes([]string{"node1"})
	if err != nil {
		t.Error("It should remove the node successfully")
	}

	if ring.Nodes.Len() != 150 {
		t.Error("it shoud have 150 nodes")
	}
}

func TestKey(t *testing.T) {
	ring := NewRing(50)
	ring.AddNodes([]string{"node::0", "node::1", "node::2"})

	keys := []string{
		"rtdd980238603650390678930",
		"rtdd980238603650390687190",
		"rtdd980238603650390691540",
		"rtdd980238603650390751690",
		"rtdd980238603650390762250",
		"rtdd980238603650390764490",
		"rtdd980238603650390767870",
		"rtdd980238603650390775870",
		"rtdd980238603650390781710",
		"rtdd980238603650390785280",
		"rtdd980238609990322235070",
		"rtdd980238609990322235980",
		"rtdd980238609990322240670",
		"rtdd980238621800396977240",
		"rtdd980238621800396981100",
		"rtdd980238621800396983000",
		"rtdd980238621800396984330",
		"rtdd980238621800396985570",
		"rtdd980238621800396985650",
		"rtdd980238621800396986560",
		"rtdd980238621800396987060",
		"rtdd980238621800397283700",
		"rtdd980238621800397286020",
		"rtdd980238621800397287840",
		"rtdd980238621800397293860",
		"rtdd980238621800397297900",
		"rtdd980238621800397298400",
		"rtdd980238621800397298650",
		"rtdd980238621800397308060",
		"rtdd980238621800397308630",
		"rtdd980238621800397341700",
		"rtdd980238621800397380800",
	}

	nodes := map[string]int{"0": 0, "1": 1, "2": 2}
	for _, key := range keys {
		node := ring.GetNode(key)
		prefixs := strings.Split(node.Id, "::")
		nodes[prefixs[len(prefixs)-1]]++
	}

	for _, v := range nodes {
		if v <= 0 {
			t.Error("It should match a node!")
			break
		}
	}
}
