package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	// === map ===
	// If the element is managed by an entity.
	mapValueIsEntity := map[int]User{0: {}}
	// Mapに入っている要素と、Mapから取得した変数は常に別物なので、下記はコンパイルエラーになる。
	// （※ mapValueIsEntity[0]としてMapから取得した時点でコピーが走るので、そもそもNameを変更しても意味がないため。）
	// mapValueIsEntity[0].Name = "Tanaka"
	mapValueIsEntity[0] = User{Name: "Tanaka"}
	fmt.Println(mapValueIsEntity[0].Name)

	// If the element is managed by pointer.
	mapValueIsPointer := map[int]*User{0: {}}
	// この場合、Mapの要素と取得した変数が同じアドレスを指すので問題ない。
	mapValueIsPointer[0].Name = "Sato"
	fmt.Println(mapValueIsPointer[0].Name)
}
