package some

import (
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

func TestMsgpack() {
	type Item struct {
		Foo string
	}

	b, err := msgpack.Marshal(&Item{Foo: "bar"})
	if err != nil {
		panic(err)
	}

	fmt.Println("msgpack Marshal bytes to string: ", string(b))

	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	fmt.Printf("msgpack Unmarshal result: %+v\n", item)
	// Output: bar
}
