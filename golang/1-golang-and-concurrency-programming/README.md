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
