# Golang, lập trình đa luồng

## Lịch sử ra đời

Golang (Go) là "đứa con" của Google, bắt đầu được thiết kế vào năm 2007 bởi bộ ba huyền thoại: Robert Griesemer, Rob Pike và Ken Thompson.

Nó được sinh ra từ sự "ức chế" của các kỹ sư Google khi phải làm việc với các hệ thống C++ khổng lồ: thời gian biên dịch quá lâu, quản lý bộ nhớ phức tạp và khó tận dụng sức mạnh của các CPU đa nhân.
Đến năm 2009, Go chính thức được mã nguồn mở và phiên bản 1.0 ra mắt vào năm 2012. Trọng tâm của Go là: Tốc độ thực thi của C++, sự an toàn của Java và tính dễ viết của Python.

## So sánh điểm mạnh/yếu của Golang với các ngôn ngữ khác

| Ngôn ngữ so sánh | Ưu điểm của Go (Go làm tốt hơn) | Nhược điểm của Go (Go làm kém hơn) |
| - | - | - |
| Python | Tốc độ thực thi nhanh hơn gấp nhiều lần; hỗ trợ đa luồng (concurrency) tự nhiên và vượt trội; phát hiện lỗi ngay lúc biên dịch (compile-time). | Hệ sinh thái thư viện AI/Machine Learning và Data Science của Go hoàn toàn lép vế so với Python; cú pháp của Go cứng nhắc hơn. |
| JavaScript (Node.js) | Có khả năng chạy đa luồng thực sự (multi-threading) thay vì bất đồng bộ trên một luồng (single-thread); hiệu năng tính toán CPU thô tốt hơn; an toàn kiểu dữ liệu (strongly typed). | JS chạy được cả ở Front-end lẫn Back-end; cộng đồng và thư viện npm của JS khổng lồ hơn rất nhiều. |
| Java | Tốn ít RAM hơn hẳn; khởi động cực nhanh (không cần máy ảo JVM); cú pháp tinh gọn, không cần viết code rườm rà (boilerplate). | Không có các framework doanh nghiệp đồ sộ và lâu đời như Spring Boot; hệ thống Lập trình hướng đối tượng (OOP) của Go rất tối giản, thiếu tính kế thừa truyền thống. |
| C++ | Có bộ thu gom rác (Garbage Collector) tự động quản lý bộ nhớ; thời gian biên dịch (compile) nhanh như chớp; cú pháp dễ học và thân thiện hơn. | C++ cho phép can thiệp cực sâu vào phần cứng và bộ nhớ (zero-overhead); C++ vẫn nhỉnh hơn về hiệu năng thuần túy trong lập trình game AAA hoặc hệ thống nhúng thời gian thực. |

## Golang được dùng để lập trình các ứng dụng như thế nào?

Nhờ đặc tính sinh ra file thực thi độc lập (binary executable), chạy nhanh, tốn ít tài nguyên và xử lý đồng thời (concurrency) cực tốt, Go thống trị trong các lĩnh vực sau:

- Hệ thống Back-end & Microservices: Go rất tuyệt vời để viết các API server chịu tải cao. Nó có thể xử lý hàng chục nghìn request mỗi giây trên các server nhỏ gọn. Các công ty như Uber, Twitch, Soundcloud đều chuyển đổi backend sang Go để tiết kiệm chi phí server.

- Công cụ DevOps & Cloud-Native: Đây là "sân nhà" của Go. Hầu hết các công cụ định hình nền tảng điện toán đám mây hiện đại đều được viết bằng Go, tiêu biểu là Docker, Kubernetes, Terraform và Prometheus.

- Hệ thống mạng & Phân tán (Network & Distributed Systems): Go hỗ trợ lập trình mạng (TCP/UDP, HTTP) cực kỳ mạnh mẽ. Nó thường được dùng để viết các hệ thống chat theo thời gian thực, hệ thống streaming video, hoặc các Load Balancer.

- Ứng dụng dòng lệnh (CLI Tools): Nhờ khả năng biên dịch ra một file chạy duy nhất cho từng hệ điều hành (Windows, macOS, Linux) mà không cần cài thêm môi trường (như Node.js hay Python), Go là lựa chọn số 1 để viết các công cụ dòng lệnh cho lập trình viên.

## Function

[Function: Input (Variadic function), Output (Multiple return), Error / Exception handling](./cmd/function/function.go)

## For loop

[For loop](./topics/for-loop/main.go)

## Package, import

[Package, import](./topics/package-import/main.go)

## Go module

Go Module là hệ thống quản lý dependency chính thức của Go, giúp quản lý thư viện và phiên bản.

### Các lệnh cơ bản

- `go mod init <tên-module>`: Khởi tạo project mới (tạo file `go.mod`).

- `go mod tidy`: Tự động thêm các module thiếu và xóa các module không còn sử dụng.

- `go get <địa-chỉ-repo>`: Tải một thư viện cụ thể (ví dụ: `go get github.com/gin-gonic/gin`).

- `go mod vendor`: Tạo thư mục `vendor/` chứa bản sao các thư viện để dùng offline.

### File `go.mod` và `go.sum`
`go.mod`: Chứa tên module, phiên bản Go và danh sách các thư viện (dependencies) cùng phiên bản của chúng.

`go.sum`: Chứa mã băm (checksum) của các thư viện để đảm bảo tính bảo mật và chắc chắn rằng code không bị thay đổi khi tải lại.

### Ví dụ dùng thư viện `rsc.io/quote`

- Chạy `go get rsc.io/quote` để tải thư viện về.
- Sử dụng thư viện trong mã nguồn.
- Chạy `go mod tidy` để cập nhật lại `go.mod` và `go.sum`.

[Go Module example](./topics/go-module/main.go)

## Naming convention

### Quy tắc Exported (Công khai) và Unexported (Nội bộ)

Đây là quy tắc quan trọng nhất của Go:

Viết hoa chữ cái đầu: Public (Exported). Có thể được truy cập từ các package khác.

Ví dụ: `func CalculateTotal()`, `type User struct`.

Viết thường chữ cái đầu: Private (Unexported). Chỉ được dùng trong nội bộ package đó.

Ví dụ: `func connectDb()`, `var secretToken string`.

### Kiểu viết chữ (MixedCaps)

Go sử dụng PascalCase (cho Exported) và camelCase (cho Unexported).

KHÔNG dùng dấu gạch dưới (_) trong tên biến (ví dụ: user_id là sai, phải là userID).

Từ viết tắt (Initialisms): Nếu tên có chứa các từ viết tắt như URL, ID, HTTP, IP, bạn phải viết hoa toàn bộ hoặc viết thường toàn bộ chúng để giữ tính nhất quán.

Đúng: `userID`, `urlRequest`, `HTTPClient`.

Sai: `userId`, `urlrequest`, `HttpClient`.

### Tên Package (Gói)

Ngắn gọn, viết thường toàn bộ, là một danh từ đơn giản.

Không dùng gạch dưới hoặc camelCase.

Tránh các tên chung chung như util, common, shared.

Đúng: `package encoding`, `package net/http`.

### Tên Interface

Nếu một Interface chỉ có một phương thức, tên của nó thường kết thúc bằng hậu tố -er.

Ví dụ: Interface có hàm `Read()` thì tên là `Reader`, hàm `Write()` thì tên là `Writer`.

Nếu có nhiều phương thức, hãy đặt một tên danh từ mô tả đúng bản chất (ví dụ: `ReadWriter`, `Car`).

### Biến Receiver (Trong Method)
Nên dùng 1 hoặc 2 chữ cái đại diện cho kiểu dữ liệu, không dùng this hay self.

`func (c *Client) Call() { ... } // Dùng 'c', không dùng 'this'`

### Tên biến ngắn (Short names)
Biến có phạm vi (scope) càng ngắn thì tên càng nên ngắn.

Trong vòng lặp: dùng `i`, `v`.

Trong hàm ngắn: dùng `r` cho Reader, `err` cho error.

Biến có phạm vi rộng hoặc biến toàn cục mới cần tên mô tả dài hơn.

### Data structure

[Data structure](./topics/data-structure/main.go)

### Multithreading

Trong Go, chúng ta không dùng khái niệm "thread" của hệ điều hành một cách trực tiếp. Thay vào đó, Go sử dụng *Goroutines* – các "luồng siêu nhẹ" được quản lý bởi Go Runtime.

Dưới đây là 3 trụ cột chính của lập trình đa luồng (concurrency) trong Go:

1. Goroutines (Luồng siêu nhẹ)

Một Goroutine chỉ tốn khoảng 2KB bộ nhớ khởi tạo (so với vài MB của Thread hệ điều hành). Bạn có thể chạy hàng triệu Goroutines trên một máy tính thông thường.

- Cách dùng: Thêm từ khóa go trước một lời gọi hàm.


``` go
func sayHello() {
    fmt.Println("Hello từ Goroutine!")
}

func main() {
    go sayHello() // Chạy hàm này ở một luồng khác
    // Lưu ý: Nếu main kết thúc, các goroutines khác cũng bị đóng ngay lập tức
}
```

2. Channels (Kênh truyền tin)

Triết lý của Go là: "Đừng giao tiếp bằng cách chia sẻ bộ nhớ; hãy chia sẻ bộ nhớ bằng cách giao tiếp." Channels được dùng để gửi và nhận dữ liệu giữa các Goroutines.

``` go
ch := make(chan string)

go func() {
    ch <- "Dữ liệu từ luồng con" // Gửi dữ liệu vào channel
}()

msg := <-ch // Nhận dữ liệu (lệnh này sẽ đợi cho đến khi có dữ liệu)
fmt.Println(msg)
```

3. Đồng bộ hóa (Sync Package)

Đôi khi bạn cần các cơ chế truyền thống để kiểm soát luồng:

- sync.WaitGroup: Đợi một nhóm Goroutines hoàn thành rồi mới chạy tiếp.
- sync.Mutex: Khóa (Lock) dữ liệu để tránh tình trạng nhiều luồng ghi vào cùng một biến một lúc (Race Condition).

So sánh Thread truyền thống vs Goroutine
|Đặc điểm|Thread (Java/C++)|Goroutine (Go)|
| - | - | - |
|Kích thước|~1-2 MB|~2 KB|
|Quản lý|Hệ điều hành (OS)|Go Runtime|
|Chuyển ngữ cảnh|Chậm (tốn tài nguyên OS)|Rất nhanh|

4. Lệnh `select`

Đây là cấu trúc điều khiển đặc biệt giúp một Goroutine có thể đợi trên nhiều thao tác channel cùng lúc.

``` go 
select {
case msg1 := <-ch1:
    fmt.Println("Nhận từ ch1:", msg1)
case ch2 <- "hi":
    fmt.Println("Gửi vào ch2 thành công")
default:
    fmt.Println("Không có gì xảy ra")
}
```
Lưu ý quan trọng: Đa luồng trong Go thiên về Concurrency (cấu trúc chương trình có thể chạy đồng thời) hơn là chỉ đơn thuần là Parallelism (chạy cùng lúc trên nhiều CPU).

### Common package

Go có một thư viện tiêu chuẩn (Standard Library) cực kỳ mạnh mẽ, giúp bạn xây dựng ứng dụng mà ít khi phải phụ thuộc vào thư viện bên ngoài.

Dưới đây là các package phổ biến nhất được chia theo nhóm chức năng:

1. Nhóm Nhập/Xuất & Chuỗi

- `fmt`: Định dạng dữ liệu (I/O). Dùng để in ra màn hình (`Println`, `Printf`) hoặc quét dữ liệu nhập vào (Scan).

- `strings`: Các hàm xử lý chuỗi như `Contains`, `Join`, `Replace`, `Split`.

- `strconv`: Chuyển đổi chuỗi thành các kiểu dữ liệu cơ bản (như `string` sang `int`) và ngược lại.

- `io` / `os`: Làm việc với hệ điều hành và các luồng dữ liệu (đọc/ghi file, biến môi trường, đối số dòng lệnh).

2. Nhóm Web & Mạng

- `net/http`: Package quan trọng nhất để xây dựng HTTP Server hoặc HTTP Client.

- `encoding/json`: Mã hóa và giải mã dữ liệu JSON (thường dùng cho API).

- `net/url`: Phân tích và xử lý các thành phần của URL.

3. Nhóm Xử lý Dữ liệu & Thời gian

- `time`: Đo lường, hiển thị và định dạng thời gian.

- `math`: Các hàm toán học cơ bản (`Abs`, `Max`, `Min`, `Sqrt`).

- `sort`: Sắp xếp các slice và tập hợp dữ liệu.

4. Nhóm Đa luồng & Đồng bộ

- `sync`: Cung cấp các công cụ đồng bộ hóa như `Mutex`, `WaitGroup`, `Once`.

- `context`: Quản lý thời hạn (deadline), tín hiệu hủy (cancellation) và các giá trị truyền qua ranh giới API/luồng.

5. Nhóm Kiểm thử & Gỡ lỗi

`testing`: Framework mặc định để viết Unit Test trong Go.

`log`: Ghi lại nhật ký hoạt động của chương trình.

`errors`: Thao tác và kiểm tra lỗi.

[Ví dụ sử dụng 1 số common package](./topics/common-package/main.go)
