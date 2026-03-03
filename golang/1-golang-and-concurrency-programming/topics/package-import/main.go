package main

import (
	"fmt"

	"example.com/golang/topics/package-import/auth" // import đường dẫn đến thư mục chứa package
)

func main() {
	ok := auth.Verify("admin")
	fmt.Println(ok, auth.ADMIN)
}
