package main

func main() {
	// struct là tập hợp các trường (fields). Đây là cách chính để tạo ra các thực thể dữ liệu phức tạp
	type User struct {
		ID    int
		Name  string
		Email string
	}

	// Array (Mảng): Có kích thước cố định, ít được dùng trực tiếp.
	// khai báo array
	_ = [3]int{1, 2, 3}
	_ = [...]int{1, 2, 3}

	// Slice (Mảng động): Linh hoạt, có thể co giãn. Đây là cấu trúc dữ liệu phổ biến nhất trong Go để lưu trữ danh sách.
	// khai báo slice
	fruits := []string{"apple", "banana"}
	fruits = append(fruits, "orange") // Thêm phần tử

	// Map (Bảng băm): Dùng để lưu trữ dữ liệu theo cặp key-value. Tương tự như Dictionary trong Python hay Object trong JS.
	// Khai báo: map[key_type]value_type
	scores := make(map[string]int)
	scores["Tuyen"] = 100

	// Interface (Giao tiếp)
	// Interface không lưu trữ dữ liệu mà định nghĩa hành vi (phương thức). Nó giúp Go đạt được tính đa hình (polymorphism).
	type Shape interface {
		Area() float64
	}
}
