Trong PostgreSQL, **Subquery (truy vấn con)** là một câu lệnh `SELECT` được lồng bên trong một câu lệnh SQL khác.

Nó có thể xuất hiện trong:

* `SELECT`
* `FROM`
* `WHERE`
* `HAVING`
* `INSERT`
* `UPDATE`
* `DELETE`

---

# Phân loại Subquery

## Scalar Subquery

Trả về **một giá trị duy nhất**.

```sql
SELECT name
FROM users
WHERE id = (
    SELECT user_id
    FROM orders
    WHERE id = 10
);
```

Điều kiện:

* Subquery phải trả về đúng 1 dòng, 1 cột.

---

## Subquery trả về nhiều dòng

Dùng với `IN`, `ANY`, `ALL`.

```sql
SELECT *
FROM users
WHERE id IN (
    SELECT user_id
    FROM orders
);
```

---

## Correlated Subquery

Subquery phụ thuộc vào query bên ngoài.

```sql
SELECT u.name
FROM users u
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.user_id = u.id
);
```

Ở đây:

* Subquery dùng `u.id` từ query ngoài.

Có thể chạy nhiều lần nếu planner không rewrite.

---

## Subquery trong FROM (Derived Table)

```sql
SELECT avg_total
FROM (
    SELECT AVG(total) AS avg_total
    FROM orders
) t;
```

Thường có thể thay bằng CTE.

---

## Subquery trong SELECT

```sql
SELECT
    name,
    (SELECT COUNT(*)
     FROM orders o
     WHERE o.user_id = u.id) AS order_count
FROM users u;
```

Đây là correlated scalar subquery.

---

# PostgreSQL xử lý Subquery thế nào?

Postgres có thể:

* Rewrite thành JOIN
* Inline subquery
* Hoặc thực thi nested loop

Planner dựa trên cost model.

---
