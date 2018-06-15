package main

func main() {
	abc := map[string]interface{}{
		"name": 3,
	}

	if value,ok := abc["name"]; ok {
		a := "123" + value.(int)
	}
}
