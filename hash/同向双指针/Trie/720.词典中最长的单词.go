/*
 * @lc app=leetcode.cn id=720 lang=golang
 * @lcpr version=21909
 *
 * [720] 词典中最长的单词
 */

// @lc code=start
package main

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		// 判断整个单词是否可以由前一个字母结尾递推过去
		if node.children[ch] == nil || !node.children[ch].isEnd {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func longestWord(words []string) (longest string) {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}
	for _, word := range words {
		if t.Search(word) && (len(word) > len(longest) || len(word) == len(longest) && word < longest) {
			longest = word
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// ["w","wo","wor","worl", "world"]\n
// @lcpr case=end

// @lcpr case=start
// ["a", "banana", "app", "appl", "ap", "apply", "apple"]\n
// @lcpr case=end

*/
