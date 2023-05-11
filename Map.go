package main

import "fmt"

func Map() {
	testMap := make(map[NamespacedName]*ProviderState)

	key1 := NamespacedName{
		Name:      "test",
		Namespace: "test1",
	}

	key2 := NamespacedName{
		Name:      "test",
		Namespace: "test1",
	}

	key3 := &NamespacedName{
		Name:      "test",
		Namespace: "test1",
	}

	fmt.Printf("%p\n", &key3)
	fmt.Println(key1 == key2)
	fmt.Printf("%p\n", &key1)
	fmt.Printf("%p\n", &key2)

	testMap[key1] = &ProviderState{
		Generation:    1,
		SentinelValue: "MySentinel",
	}

	testMap[key2] = &ProviderState{
		Generation:    3,
		SentinelValue: "MySentinel",
	}

	fmt.Println(testMap[key1].Generation)
}

type NamespacedName struct {
	Name      string
	Namespace string
}

type ProviderState struct {
	Generation    int64 //TODO, suppose to use string
	SentinelValue string
}
