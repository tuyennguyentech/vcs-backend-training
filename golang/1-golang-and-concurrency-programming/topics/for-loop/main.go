package main

import "fmt"

func main() {

	// trong go,

	i := 1
	for i <= 3 { // vòng lặp theo kiểu while trong c++
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ { // vòng lặp theo kiểu for trong C++
		fmt.Println(j)
	}

	for i := range 3 { // thực hiện vòng lặp 3 lần
		fmt.Println("range", i)
	}

	for { // vòng lặp vô hạn
		fmt.Println("loop")
		break // dùng break để thoát vòng lặp
	}

	for n := range 6 {
		if n%2 == 0 {
			continue // dùng continue để nhảy qua lần lặp kế tiếp
		}
		fmt.Println(n)
	}

	nums := []int{2, 3, 4}
	sum := 0
	// lặp qua các phần tử trong slice, tương tự với array
	// bỏ qua index bằng cách dùng blank identifier _
	// num là giá trị phần tử trong slice
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// dùng cả index và giá trị phần tử
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	// duyệt qua tất cả các cặp key-value trong map
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// chỉ duyệt qua key
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// duyệt qua từng kí tự Unicode string kèm với chỉ số thứ tự của nó
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
