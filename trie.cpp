#include <bits/stdc++.h>
using namespace std;

// A Trie node
struct Trie
{
	bool isLeaf;
	unordered_map<char, Trie*> map;
};

// Function that returns a new Trie node
Trie* getNewTrieNode()
{
	Trie* node = new Trie;
	node->isLeaf = false;

	return node;
}

// Iterative function to insert a string in Trie.
void insert(Trie*& head, char* str)
{
	if (head == nullptr)
        head = getNewTrieNode();
	Trie* curr = head;
	while (*str)
	{
		if (curr->map.find(*str) == curr->map.end())
			curr->map[*str] = getNewTrieNode();

		curr = curr->map[*str];
		str++;
	}
	curr->isLeaf = true;
}

// returns true if given node has any children
bool haveChildren(Trie const* curr)
{
	for (auto it : curr->map)
		if (it.second != nullptr)
			return true;
	return false;
}

// Recursive function to delete a string in Trie.
bool deletion(Trie*& curr, char* str)
{
	if (curr == nullptr)
		return false;
	if (*str)
	{
		if (curr != nullptr &&  curr->map.find(*str) != curr->map.end() &&
			deletion(curr->map[*str], str + 1) && curr->isLeaf == false)
		{
		    if (!haveChildren(curr))
			{
				delete curr;;
				curr = nullptr;
				return true;
			}
			else {
				return false;
			}
		}
	}
	if (*str == '\0' && curr->isLeaf)
	{
		if (!haveChildren(curr))
		{
			delete curr;;
			curr = nullptr;
			return true;
		}
		else
		{
			curr->isLeaf = false;
			return false;
		}
	}

	return false;
}

// Iterative function to search a string in Trie
bool search(Trie* head, char* str)
{
	if (head == nullptr)
		return false;

	Trie* curr = head;
	while (*str)
	{
		curr = curr->map[*str];
		if (curr == nullptr)
			return false;
		str++;
	}
	return curr->isLeaf;
}

int main()
{
	Trie* head = nullptr;
	insert(head, "hello");
	cout << search(head, "hello") << " ";   	
    // print 1
	deletion(head, "hello");
	cout << search(head, "hello") << " ";   	
    // print 0 (hello deleted)		
	return 0;
}