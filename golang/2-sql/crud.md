Trong PostgreSQL, CRUD là bốn thao tác cơ bản trên dữ liệu:

* **C**reate → `INSERT`
* **R**ead → `SELECT`
* **U**pdate → `UPDATE`
* **D**elete → `DELETE`


---

# CREATE – INSERT

## Tạo bảng

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE,
    created_at TIMESTAMP DEFAULT now()
);
```

---

## Chèn dữ liệu

### Insert 1 record

```sql
INSERT INTO users (name, email)
VALUES ('Tuyen', 'tuyen@example.com');
```

### Insert nhiều record

```sql
INSERT INTO users (name, email)
VALUES
    ('Alice', 'alice@example.com'),
    ('Bob', 'bob@example.com');
```

---

## INSERT + RETURNING (đặc trưng mạnh của Postgres)

```sql
INSERT INTO users (name, email)
VALUES ('Charlie', 'charlie@example.com')
RETURNING id;
```

Postgres cho phép lấy ngay row vừa insert.

---

## UPSERT (INSERT ... ON CONFLICT)

```sql
INSERT INTO users (email, name)
VALUES ('alice@example.com', 'Alice Updated')
ON CONFLICT (email)
DO UPDATE SET name = EXCLUDED.name;
```

---

# READ – SELECT

## Đọc toàn bộ

```sql
SELECT * FROM users;
```

---

## Có điều kiện

```sql
SELECT id, name
FROM users
WHERE email = 'alice@example.com';
```

---

## ORDER + LIMIT

```sql
SELECT *
FROM users
ORDER BY created_at DESC
LIMIT 10;
```

---

## JOIN

```sql
SELECT u.name, o.total
FROM users u
JOIN orders o ON o.user_id = u.id;
```

---

## Transaction snapshot (MVCC)

Postgres dùng MVCC:

* SELECT không block UPDATE
* UPDATE không block SELECT
* Mỗi transaction thấy snapshot riêng

---

# UPDATE

## Update cơ bản

```sql
UPDATE users
SET name = 'Alice Smith'
WHERE id = 1;
```

---

## Update + RETURNING

```sql
UPDATE users
SET name = 'Alice Final'
WHERE id = 1
RETURNING *;
```

---

## Lưu ý quan trọng (MVCC)

Trong Postgres:

* UPDATE không overwrite row
* Nó tạo version mới
* Row cũ được đánh dấu expired (xmin/xmax)

→ Cần VACUUM để dọn dẹp.

---

# DELETE

## Xóa có điều kiện

```sql
DELETE FROM users
WHERE id = 5;
```

---

## DELETE + RETURNING

```sql
DELETE FROM users
WHERE id = 5
RETURNING *;
```

---

## DELETE trong Postgres thực chất:

* Không xóa ngay vật lý
* Chỉ mark tuple là dead
* VACUUM sẽ reclaim space

---

# CRUD trong Transaction

```sql
BEGIN;

INSERT INTO users (name) VALUES ('Test');

UPDATE users SET name = 'Updated' WHERE id = 10;

DELETE FROM users WHERE id = 11;

COMMIT;
```

Hoặc rollback:

```sql
ROLLBACK;
```

---

# CRUD và Locking Behavior

Postgres:

| Operation | Lock type        |
| --------- | ---------------- |
| SELECT    | AccessShareLock  |
| INSERT    | RowExclusiveLock |
| UPDATE    | RowExclusiveLock |
| DELETE    | RowExclusiveLock |

Nhưng:

* Row-level locking, không lock toàn table
* MVCC giảm deadlock mạnh

---

# CRUD nâng cao

## Soft delete

```sql
UPDATE users
SET deleted_at = now()
WHERE id = 1;
```

---

## Batch update

```sql
UPDATE users
SET active = false
WHERE last_login < now() - interval '1 year';
```

---

## Pagination chuẩn

```sql
SELECT *
FROM users
ORDER BY id
LIMIT 20 OFFSET 40;
```

Hoặc tốt hơn (keyset pagination):

```sql
SELECT *
FROM users
WHERE id > 40
ORDER BY id
LIMIT 20;
```

---
