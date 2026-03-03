# SQL

## PostgreSQL overview

PostgreSQL là 1 object-relational database management system (ORDBMS) mã nguồn mở. Nó hỗ trợ phần lớn tiêu chuẩn SQL và cung cấp nhiều tính năng hiện đại như:

- complex queries
- foreign keys
- updatable views
- transactional integrity
- multiversion concurrency control

Ngoài ra PostgreSQL có thể được mở rộng bởi người dùng bằng nhiều cách, ví dụ như thêm mới:

- data types
- functions
- operators
- aggregate functions
- index methods
- procedural languages

## So sánh điểm mạnh, điểm yếu với các SQL Database khác



|Danh mục|MySQL|PostgreSQL|
| - | - | - |
|Công nghệ cơ sở dữ liệu|MySQL đơn thuần là một hệ thống quản lý cơ sở dữ liệu quan hệ.|PostgreSQL là một hệ thống quản lý cơ sở dữ liệu quan hệ đối tượng.|
|Tính năng|MySQL hỗ trợ số ít các tính năng cơ sở dữ liệu như chế độ xem, điều kiện kích hoạt và quy trình.|PostgreSQL hỗ trợ hầu hết các tính năng cơ sở dữ liệu nâng cao như chế độ xem cụ thể hóa, điều kiện kích hoạt INSTEAD OF và quy trình được lưu trữ bằng nhiều ngôn ngữ.|
|Loại dữ liệu|MySQL hỗ trợ các loại dữ liệu số, ký tự, ngày và giờ, không gian và JSON.|PostgreSQL hỗ trợ tất cả các loại dữ liệu MySQL cùng với các loại dữ liệu hình học, liệt kê, địa chỉ mạng, mảng, phạm vi, XML, hstore và kết hợp.|
|Tuân thủ ACID|MySQL chỉ tuân thủ ACID với công cụ lưu trữ InnoDB và NDB Cluster.|PostgreSQL luôn tuân thủ ACID.|
|Chỉ mục|MySQL hỗ trợ chỉ mục B-tree và R-tree.|PostgreSQL hỗ trợ nhiều loại chỉ mục như chỉ mục biểu thức, chỉ mục một phần và chỉ mục băm cùng với dạng cây.|
|Hiệu năng|MySQL có hiệu năng cao hơn đối với các thao tác đọc thường xuyên.|PostgreSQL có hiệu năng cao hơn đối với các thao tác ghi thường xuyên.|
|Hỗ trợ người mới bắt đầu|MySQL dễ hơn khi bắt đầu sử dụng. MySQL có một bộ công cụ đa dạng hơn cho người dùng không chuyên về kỹ thuật.|PostgreSQL phức tạp hơn khi bắt đầu sử dụng. PostgreSQL có một bộ công cụ hạn chế cho người dùng không chuyên về kỹ thuật.|

|Danh mục|Oracle|PostgreSQL|
| - | - | - |
| **Giấy phép và chi phí**                   | Mô hình cấp phép thương mại, yêu cầu trả phí bản quyền        | Mã nguồn mở và miễn phí sử dụng                                             |
| **Cân nhắc chi phí**                       | Chi phí ban đầu và duy trì cao cho bản quyền và hỗ trợ        | Chi phí thấp hơn, nhưng có thể phát sinh phí hỗ trợ doanh nghiệp            |
| **Hiệu năng và khả năng mở rộng**          | Hiệu năng cao, tối ưu cho khả năng mở rộng cấp doanh nghiệp   | Hiệu năng mạnh, hỗ trợ mở rộng theo chiều ngang và chiều dọc                |
| **Benchmark**                              | Được sử dụng trong các doanh nghiệp lớn với workload khắt khe | Liên tục cải thiện hiệu năng qua từng phiên bản                             |
| **Bảo mật và tuân thủ**                    | Tính năng bảo mật nâng cao (Mã hóa dữ liệu, Label Security)   | Bảo mật tích hợp sẵn (RLS, SSL, TLS)                                        |
| **Hỗ trợ tuân thủ**                        | Tuân thủ mạnh các tiêu chuẩn như HIPAA, GDPR, SOC 2           | Hỗ trợ tuân thủ nhưng cần cấu hình bổ sung                                  |
| **Khả năng mở rộng và tùy biến**           | Extension PL/SQL độc quyền                                    | Extension mã nguồn mở và hàm tùy chỉnh                                      |
| **Định dạng hỗ trợ**                       | Hỗ trợ JSON, XML và các ngôn ngữ thủ tục                      | Hỗ trợ mạnh JSON, XML và ngôn ngữ thủ tục                                   |
| **Tuân thủ SQL**                           | Mức độ tuân thủ cao nhưng có nhiều mở rộng SQL độc quyền      | Tuân thủ cao tiêu chuẩn SQL, theo nguyên tắc mã nguồn mở                    |
| **Tương thích và replication**             | Hỗ trợ đa nền tảng, Oracle GoldenGate, Active Data Guard      | Hỗ trợ đa nền tảng, streaming replication và logical replication            |
| **Tính năng**                              | Index nâng cao, partitioning và sharding mạnh                 | Hỗ trợ indexing, partitioning và sharding thông qua công cụ mã nguồn mở     |
| **Công cụ tích hợp sẵn vs bên thứ ba**     | Công cụ tối ưu và tuning tích hợp sẵn                         | Nhiều công cụ bên thứ ba cho tuning hiệu năng                               |
| **High Availability và Disaster Recovery** | Oracle RAC, Data Guard, Flashback hỗ trợ HA/DR mạnh           | Streaming replication, logical replication và cơ chế failover               |
| **Backup và khôi phục**                    | Tùy chọn backup/restore nâng cao, công cụ tự động hóa         | Công cụ backup/restore tin cậy nhưng có thể cần cấu hình thủ công nhiều hơn |
| **Hệ sinh thái và cộng đồng**              | Hỗ trợ mạnh từ vendor, quan hệ đối tác doanh nghiệp           | Cộng đồng mã nguồn mở lớn và đang mở rộng adoption doanh nghiệp             |
| **Tích hợp bên thứ ba**                    | Tích hợp sâu với phần mềm doanh nghiệp                        | Hệ sinh thái công cụ bên thứ ba mạnh                                        |

|SQL Server|PostgreSQL|
| - | - |
|Relational database management system|Object-relational database management system|
|Commercial product from Microsoft|Open source (completely free)|
|Runs only on Microsoft or Linux|Runs on most machines and operating systems|
|Uses Transact-SQL or T-SQL (standard SQL + extra functionality)|Uses Standard SQL|

|Hệ quản trị|Ưu điểm|Nhược điểm|
| - | - | - |
| **PostgreSQL** | - Khả năng mở rộng cao (có thể thêm function, data type, procedural language, extension…)  <br> - Hỗ trợ dữ liệu phi cấu trúc (audio, video, hình ảnh, JSON…)  <br> - MVCC cho xử lý đồng thời tốt, hỗ trợ lượng giao dịch cao với rất ít deadlock  <br> - Hỗ trợ high availability và khôi phục khi server lỗi  <br> - Tính năng bảo mật nâng cao (mã hóa dữ liệu, SSL, xác thực nâng cao…)  <br> - Cộng đồng mã nguồn mở lớn, liên tục cải tiến | - Có thể chậm hơn SQL Server hoặc MySQL trong một số workload cụ thể  <br> - Tập trung nhiều vào tính tương thích tiêu chuẩn SQL, nên tối ưu tốc độ đôi khi cần cấu hình thêm  <br> - Cài đặt và cấu hình ban đầu có thể khó với người mới |
| **SQL Server** | - Hiệu năng cao, hỗ trợ in-memory database  <br> - Nhiều tính năng bảo mật tích hợp (cảnh báo, monitoring, bảo vệ và phân loại dữ liệu…)  <br> - Cài đặt và cấu hình đơn giản, giao diện dễ dùng, cập nhật tự động  <br> - Backup và recovery thuận tiện, công cụ HA mạnh  <br> - Có thể lập lịch tác vụ qua SQL Server Management Studio  <br> - Tích hợp tốt với hệ sinh thái Microsoft (analytics, development, monitoring tools…)             | - Không sử dụng MVCC mặc định, phụ thuộc vào cơ chế locking để đảm bảo nhất quán  <br> - Chi phí bản quyền, hỗ trợ và tính năng nâng cao cao  <br> - Có thể yêu cầu nâng cấp phần cứng để chạy các phiên bản mới|
