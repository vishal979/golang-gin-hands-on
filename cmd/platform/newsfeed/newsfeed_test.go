package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{})

	if len(feed.Items) != 1 {
		t.Errorf("Item not added")
	}

}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	results := feed.GetAll()
	if len(results) != 1 {
		t.Errorf("item not added")
	}
}
