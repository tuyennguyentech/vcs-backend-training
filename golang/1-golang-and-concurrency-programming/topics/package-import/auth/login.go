package auth // convention tên package ko dùng kí tự đặc biệt, không viết hoa, ngắn gọn, thường đặt theo tên thư mục

// viết hoa chữ cái đầu của tên biến hoặt tên hàm để có thể được export ra ngoài package
func Verify(user string) bool {
	return user == ADMIN
}
