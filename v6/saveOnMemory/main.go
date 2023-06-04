package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	/*
		注釈
		このコードでは、PostByIdとPostsByAuthorをvarで宣言した後、
		make関数を使用してそれぞれのマップを初期化しています。
		Golangにおいて、マップは参照型です。
		変数を宣言するだけではマップ自体は初期化されず、nilとして初期化されます。
		nilマップはメモリを割り当てていないため、
		要素を追加しようとするとランタイムエラーが発生します。
		そのため、make関数を使用してマップを初期化する必要があります。
		make関数は、指定された型の初期化された値を作成し、そのポインタを返します。
		この場合、make(map[int]*Post)はPostByIdとPostsByAuthorの型に合わせた初期化されたマップを作成し、
		そのポインタを変数に代入することで、有効なマップを作成しています。
		したがって、var PostById map[int]*Postとvar PostsByAuthor map[string][]*Postで変数を宣言し、
		make(map[int]*Post)とmake(map[string][]*Post)で初期化しているのは、
		マップを使用するための正しい形式です。
	*/
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "愛はあるんか", Author: "俺"}
	post2 := Post{Id: 2, Content: "愛はあるんか2", Author: "俺2"}
	post3 := Post{Id: 3, Content: "愛はあるんか3", Author: "俺3"}

	store(post1)
	store(post2)
	store(post3)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["俺"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["俺2"] {
		fmt.Println(post)

	}
}
