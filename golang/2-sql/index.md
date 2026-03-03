# PostgreSQL Indexes


---

## Giới thiệu về PostgreSQL Index

Trong PostgreSQL, index là một cấu trúc dữ liệu giúp tăng tốc độ truy xuất dữ liệu. Nó cung cấp cho PostgreSQL một cách nhanh chóng để xác định vị trí các dòng trong bảng mà không cần quét toàn bộ bảng.

Có thể hình dung index giống như mục lục của một cuốn sách. Thay vì đọc từng trang để tìm một chủ đề, chỉ cần xem mục lục để biết chính xác trang cần tìm. PostgreSQL sử dụng index theo cách tương tự để tìm dữ liệu nhanh hơn.

Giả sử có bảng `contacts` với cấu trúc như sau:

```sql
CREATE TABLE contacts (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(10) NOT NULL
);
```

Thực hiện truy vấn sau để tìm một liên hệ theo tên:

```sql
SELECT * FROM contacts
WHERE name = 'John Doe';
```

Nếu không có index, PostgreSQL phải quét toàn bộ bảng `contacts` để tìm `"John Doe"`. Nếu bảng chứa nhiều dòng, thao tác này sẽ chậm — tương tự như đọc cả cuốn sách chỉ để tìm một dòng duy nhất.

Nếu có index trên cột `name`, PostgreSQL có thể định vị các dòng phù hợp nhanh hơn nhiều.

---

## Tạo Index

Để tạo index, sử dụng câu lệnh `CREATE INDEX`:

```sql
CREATE INDEX contacts_name
ON contacts(name);
```

Câu lệnh này tạo một index có tên `contacts_name` trên cột `name` của bảng `contacts`.

Sau khi index được tạo, PostgreSQL trích xuất toàn bộ giá trị từ cột `name` và lưu chúng vào cấu trúc dữ liệu của index. Quá trình này có thể mất thời gian nếu bảng chứa nhiều dòng — giống như việc tạo mục lục cho một cuốn sách dày sẽ lâu hơn.

Theo mặc định:

* Cho phép các thao tác `SELECT` trong quá trình tạo index.
* Chặn các thao tác `INSERT`, `UPDATE`, và `DELETE` để đảm bảo an toàn dữ liệu.

Nếu không thể chặn các thao tác ghi trong quá trình tạo index, có thể sử dụng cú pháp:

```sql
CREATE INDEX CONCURRENTLY contacts_name
ON contacts(name);
```

Tùy chọn `CONCURRENTLY` cho phép tạo index mà không khóa các thao tác `INSERT`, `UPDATE`, hoặc `DELETE`, nhưng tốc độ tạo sẽ chậm hơn đáng kể.

---

## Sử dụng Index

Khi thực hiện truy vấn:

```sql
SELECT * FROM contacts
WHERE name = 'John Doe';
```

PostgreSQL có thể sử dụng index `contacts_name` để nhanh chóng tìm các dòng phù hợp.

Sau khi index được tạo, PostgreSQL phải giữ index đồng bộ với bảng.

Ví dụ:

* Khi thêm dòng mới vào bảng `contacts`, index cũng được cập nhật.
* Khi cập nhật hoặc xóa dòng, index cũng được cập nhật tương ứng.

Vì vậy:

* Index giúp tăng tốc độ đọc dữ liệu.
* Đồng thời làm tăng một phần chi phí cho các thao tác ghi (`INSERT`, `UPDATE`, `DELETE`).

Điều này tương tự như khi chỉnh sửa nội dung sách thì cũng phải cập nhật lại mục lục.

---

# Các loại Index trong PostgreSQL

PostgreSQL cung cấp nhiều loại index khác nhau, mỗi loại phù hợp với những tình huống dữ liệu và mẫu truy vấn cụ thể.

Hiểu rõ các loại index giúp tối ưu hiệu năng truy vấn hiệu quả hơn.

---

## B-tree Index

B-tree là loại index mặc định trong PostgreSQL. B-tree là viết tắt của “balanced tree”.

B-tree lưu trữ giá trị theo thứ tự sắp xếp, do đó rất hiệu quả cho:

* So sánh bằng (`=`)
* So sánh phạm vi (`>`, `<`, `>=`, `<=`)
* Truy vấn khoảng (`BETWEEN`)
* Sắp xếp (`ORDER BY`)

---

## Hash Index

Hash index lưu trữ mã băm 32-bit được tạo từ giá trị của các cột được index.

Vì vậy, hash index chỉ xử lý được các so sánh bằng (`=`).

Không hỗ trợ truy vấn phạm vi hoặc sắp xếp.

---

## GIN Index

GIN (Generalized Inverted Index) là index dạng đảo, phù hợp với các giá trị tổng hợp như:

* Mảng (arrays)
* Dữ liệu JSONB
* Full-text search

GIN lưu một entry riêng cho mỗi thành phần trong giá trị tổng hợp, vì vậy có thể xử lý các truy vấn kiểm tra sự tồn tại của một thành phần cụ thể.

---

## GiST Index

GiST (Generalized Search Tree) là loại index linh hoạt, hỗ trợ nhiều kiểu dữ liệu khác nhau, bao gồm:

* Dữ liệu hình học
* Full-text search

GiST cho phép nhiều chiến lược tìm kiếm như:

* Tìm phần tử gần nhất (nearest neighbor)
* Tìm kiếm khớp một phần (partial match)

Phù hợp cho các ứng dụng chuyên biệt như dữ liệu không gian.

---

## SP-GiST Index

SP-GiST (Space-Partitioned GiST) phù hợp để index dữ liệu có cấu trúc phân cấp hoặc kiểu dữ liệu phức tạp.

SP-GiST chia không gian index thành các vùng không chồng lấn, giúp tìm kiếm hiệu quả trong các cấu trúc dữ liệu đặc thù.

---

## BRIN (Block Range Index)

BRIN được thiết kế cho các bảng rất lớn, nơi việc index từng dòng là không thực tế.

BRIN chia bảng thành các phạm vi block (page ranges) và lưu thông tin tóm tắt cho mỗi phạm vi.

Nhờ vậy:

* Phù hợp cho truy vấn phạm vi trên tập dữ liệu lớn.
* Sử dụng rất ít không gian lưu trữ so với các loại index khác.
* Hiệu quả đặc biệt khi dữ liệu có tính tuần tự (ví dụ: timestamp tăng dần).
