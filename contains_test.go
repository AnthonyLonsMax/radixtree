package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestContains(t *testing.T) {

	t.Parallel()

	var tree radixtree.RadixTree
	tree.Add("worderland")
	tree.Add("word")
	tree.Add("worddy")
	tree.Add("work")
	tree.Add("worry")
	tree.Add("worries")
	tree.Add("wallet")
	tree.Add("love")
	tree.Add("lonnly")
	tree.Add("lovers")
	tree.Add("anthony")
	tree.Add("anth")

	if !tree.Contains("wallet") {
		t.Fatal("tree element should be in the tree")
	}
}
