package main

type ShoppingItem struct {
	Id   int
	Name string
}

type ShoppingList struct {
	Items map[int]ShoppingItem
}

func NewShoppingList() *ShoppingList {
	return &ShoppingList{
		Items: map[int]ShoppingItem{},
	}
}

func (s *ShoppingList) Add(item ShoppingItem) int {
	key := len(s.Items) + 1
	s.Items[key] = item
	return key
}

func (s *ShoppingList) Get(id int) ShoppingItem {
	return s.Items[id]
}

func (s *ShoppingList) List() map[int]ShoppingItem {
	return s.Items
}
