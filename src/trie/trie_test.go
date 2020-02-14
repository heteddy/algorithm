/*
@Copyright:
*/
/*
@Time : 2020/2/14 14:41
@Author : teddy
@File : trie_test.go
*/

package trie

import "testing"

func TestTrie_Add(t *testing.T) {
	trie := NewTrie()
	trie.Add("hello")
	trie.Add("how")
	if trie.ContainPrefix("ho") == 1 {
	} else {
		t.Error("he count error")
	}
	trie.Add("hedetao")
	if trie.ContainPrefix("he") == 2 {
	} else {
		t.Error("he count error")
	}
	trie.Add("heteddy")
	if trie.ContainPrefix("he") == 3 {
	} else {
		t.Error("he count error")
	}
}

func TestTrie_ContainWord(t *testing.T) {
	trie := NewTrie()
	trie.Add("hello")
	trie.Add("how")
	if trie.ContainWord("ho") {
		t.Error("he count error")
	}
	trie.Add("hedetao")
	if trie.ContainWord("hedetao") {
	} else {
		t.Error("he count error")
	}
	trie.Add("heteddy")
	if trie.ContainWord("heteddy") {
	} else {
		t.Error("he count error")
	}
}

func TestTrie_Remove(t *testing.T) {
	trie := NewTrie()
	trie.Add("hello")
	trie.Add("how")
	if trie.ContainWord("ho") {
		t.Error("he count error")
	}
	trie.Add("hedetao")
	if trie.ContainWord("hedetao") {
	} else {
		t.Error("he count error")
	}
	trie.Add("heteddy")
	if trie.ContainWord("heteddy") {
	} else {
		t.Error("he count error")
	}

	trie.Remove("how")
	if trie.ContainWord("how") {
		t.Error("he count error")
	}
}
