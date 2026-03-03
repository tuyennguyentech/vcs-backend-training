Trong PostgreSQL, **JOIN** dùng để kết hợp dữ liệu từ nhiều bảng dựa trên điều kiện liên kết (thường là khóa chính – khóa ngoại).

---

# Ví dụ cơ bản

Giả sử:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    total NUMERIC
);
```

---

# INNER JOIN

Chỉ trả về các dòng có match ở cả hai bảng.

```sql id="j1a8p3"
SELECT u.name, o.total
FROM users u
INNER JOIN orders o
    ON u.id = o.user_id;
```

Nếu user không có order → không xuất hiện.

---

# LEFT JOIN

Giữ toàn bộ bảng bên trái, kể cả khi không match.

```sql id="k4z9x1"
SELECT u.name, o.total
FROM users u
LEFT JOIN orders o
    ON u.id = o.user_id;
```

Nếu user chưa có order → `o.total = NULL`

---

# RIGHT JOIN

Ngược lại với LEFT JOIN.

```sql id="t9v6c2"
SELECT u.name, o.total
FROM users u
RIGHT JOIN orders o
    ON u.id = o.user_id;
```

---

# FULL JOIN

Giữ tất cả dòng từ cả hai bảng.

```sql id="m8w2r7"
SELECT u.name, o.total
FROM users u
FULL JOIN orders o
    ON u.id = o.user_id;
```

---

# CROSS JOIN

Tích Descartes (mọi cặp kết hợp).

```sql id="x2b4y9"
SELECT *
FROM users
CROSS JOIN orders;
```

Có thể tạo số dòng rất lớn.

---

# SELF JOIN

Join một bảng với chính nó.

```sql id="p6n3q8"
SELECT e.name, m.name AS manager
FROM employees e
LEFT JOIN employees m
    ON e.manager_id = m.id;
```

---

# JOIN nhiều bảng

```sql id="f3u1k5"
SELECT u.name, o.total, p.product_name
FROM users u
JOIN orders o ON u.id = o.user_id
JOIN products p ON o.product_id = p.id;
```

---

# Cách PostgreSQL thực thi JOIN

Postgres có 3 chiến lược chính:

| Join type   | Khi dùng                            |
| ----------- | ----------------------------------- |
| Nested Loop | Dataset nhỏ / có index              |
| Hash Join   | Không có index / dataset trung bình |
| Merge Join  | Cả hai phía đã sort                 |

Bạn có thể xem bằng:

```sql
EXPLAIN ANALYZE SELECT ...
```

---

# Hiệu năng JOIN phụ thuộc vào

1. Index đúng cột join
2. Thứ tự bảng
3. Kích thước bảng
4. work_mem
5. Statistics (ANALYZE)

---

# Best practice

* Luôn index cột dùng trong `ON`
* Tránh SELECT *
* Tránh join bảng quá lớn không cần thiết
* Kiểm tra execution plan

---

# Ví dụ tối ưu

Có index:

```sql
CREATE INDEX idx_orders_user_id ON orders(user_id);
```

Giúp join nhanh hơn rất nhiều.

---

# Tóm tắt

| Loại JOIN | Giữ bảng nào  |
| --------- | ------------- |
| INNER     | Chỉ match     |
| LEFT      | Giữ bảng trái |
| RIGHT     | Giữ bảng phải |
| FULL      | Giữ cả hai    |
| CROSS     | Mọi kết hợp   |

---
