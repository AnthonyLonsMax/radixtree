package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestReadAllTheKeys(t *testing.T) {

	t.Parallel()

	var tree radixtree.RadixTree
	tree.Add("worderland")
	tree.Add("word")
	tree.Add("worddy")
	tree.Add("work")
	tree.Add("worry")
	tree.Add("wor")
	tree.Add("worries")
	tree.Add("wallet")
	tree.Add("love")
	tree.Add("lonnly")
	tree.Add("lovers")
	tree.Add("anthony")
	tree.Add("ony")
	tree.Add("anth")

	for _, word := range *tree.Keys() {
		t.Log(word)
	}

}
