package main

import "testing"

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}

	if post.Id != 1 {
		t.Error("1じゃない")
	}
	if post.Content != "愛はあるんか" {
		t.Error("愛がない")
	}
}

func TestEncode(t *testing.T) {
	t.Skip("テストしません encode")
}
