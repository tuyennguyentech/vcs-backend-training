Trong PostgreSQL, **Foreign Key (khóa ngoại)** là một ràng buộc (constraint) dùng để đảm bảo **toàn vẹn tham chiếu (referential integrity)** giữa hai bảng.

Nói ngắn gọn:

> Foreign key đảm bảo giá trị ở bảng con phải tồn tại ở bảng cha.

---

# Khái niệm cơ bản

Giả sử có hai bảng:

* `users` (bảng cha – parent)
* `orders` (bảng con – child)

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT
);
```

```sql
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

Ở đây:

* `orders.user_id` là **foreign key**
* Nó tham chiếu đến `users.id`

---

# Ý nghĩa hoạt động

Foreign key đảm bảo:

* Không thể insert `user_id` không tồn tại trong `users`
* Không thể xóa user nếu còn order tham chiếu (trừ khi có rule khác)

---

# Ví dụ lỗi

```sql
INSERT INTO orders (user_id) VALUES (999);
```

Nếu `users.id = 999` không tồn tại → PostgreSQL báo lỗi.

---

# ON DELETE / ON UPDATE

Foreign key có thể định nghĩa hành vi khi:

* Bản ghi cha bị xóa
* Hoặc khóa chính bị thay đổi

---

## ON DELETE CASCADE

```sql
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE
```

Khi xóa user → xóa luôn orders liên quan.

---

## ON DELETE SET NULL

```sql
ON DELETE SET NULL
```

Khi xóa user → `user_id` trong orders = NULL.

---

## ON DELETE RESTRICT (mặc định)

Không cho xóa nếu còn bản ghi tham chiếu.

---

# Composite Foreign Key

Có thể tham chiếu nhiều cột:

```sql
FOREIGN KEY (order_id, product_id)
REFERENCES order_items(order_id, product_id)
```

---

# Cách thêm foreign key sau khi tạo bảng

```sql
ALTER TABLE orders
ADD CONSTRAINT fk_user
FOREIGN KEY (user_id)
REFERENCES users(id);
```

---

# Cách xóa foreign key

```sql
ALTER TABLE orders
DROP CONSTRAINT fk_user;
```

---

# Performance lưu ý quan trọng

Postgres **không tự động tạo index cho foreign key**.

Bạn nên tạo index cho cột foreign key:

```sql
CREATE INDEX idx_orders_user_id ON orders(user_id);
```

Nếu không:

* DELETE từ bảng cha có thể rất chậm
* UPDATE key cha có thể scan toàn bảng con

---

# Locking behavior

Khi insert vào bảng con:

* PostgreSQL phải kiểm tra tồn tại ở bảng cha
* Có thể acquire row-level lock

Trong workload ghi nhiều với nhiều foreign key:

* Có thể phát sinh contention

---

# Internal behavior (MVCC)

Foreign key trong Postgres được thực thi bằng:

* Trigger nội bộ
* Check tại thời điểm INSERT/UPDATE/DELETE
* Có thể deferred

---

# DEFERRABLE constraint

```sql
FOREIGN KEY (user_id)
REFERENCES users(id)
DEFERRABLE INITIALLY DEFERRED
```

→ Check constraint khi COMMIT thay vì ngay lúc insert.

Rất hữu ích trong batch transaction phức tạp.

---

# Tóm tắt

| Thành phần   | Vai trò                  |
| ------------ | ------------------------ |
| Parent table | Chứa khóa chính          |
| Child table  | Chứa foreign key         |
| Mục đích     | Đảm bảo toàn vẹn dữ liệu |
| ON DELETE    | Quy định hành vi khi xóa |
| DEFERRABLE   | Hoãn kiểm tra constraint |

---

