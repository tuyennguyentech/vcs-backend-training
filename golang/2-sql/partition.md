# 📝 Ghi chú: PostgreSQL Table Partitioning

---

## 1. Tại sao cần dùng Partition?

- **Tăng hiệu suất truy vấn** → query chỉ quét phần dữ liệu liên quan thay vì toàn bảng
- **Xóa dữ liệu cũ nhanh** → `DROP TABLE partition` nhanh hơn `DELETE` hàng triệu row
- **Sequential scan hiệu quả hơn** → khi truy vấn 1 lượng lớn trong 1 phân vùng
- **Dữ liệu ít dùng** → có thể chuyển sang storage rẻ hơn

> ⚠️ **Nguyên tắc thực tiễn:** Chỉ nên dùng khi bảng lớn hơn RAM của server

---

## 2. Các loại Partition

| Loại | Cách hoạt động | Dùng khi nào |
|------|---------------|--------------|
| **Range** | Chia theo khoảng giá trị (vd: ngày tháng) | Dữ liệu time-series, logs |
| **List** | Liệt kê rõ giá trị vào từng partition | Phân loại cố định (region, status) |
| **Hash** | Dùng modulus + remainder | Phân tán đều khi không có tiêu chí rõ ràng |

---

## 3. Declarative Partitioning (Cách hiện đại, khuyên dùng)

### Tạo bảng cha
```sql
CREATE TABLE measurement (
    city_id   int not null,
    logdate   date not null,
    peaktemp  int,
    unitsales int
) PARTITION BY RANGE (logdate);
```

### Tạo partition con
```sql
CREATE TABLE measurement_y2008m01 PARTITION OF measurement
    FOR VALUES FROM ('2008-01-01') TO ('2008-02-01');
```
> ⚠️ Giới hạn trên là **exclusive** (không bao gồm), giới hạn dưới là **inclusive**

### Tạo index
```sql
-- Tạo index trên bảng cha → tự động tạo trên tất cả partition
CREATE INDEX ON measurement (logdate);
```

### Sub-partitioning
```sql
CREATE TABLE measurement_y2006m02 PARTITION OF measurement
    FOR VALUES FROM ('2006-02-01') TO ('2006-03-01')
    PARTITION BY RANGE (peaktemp);  -- phân vùng thêm theo cột khác
```

---

## 4. Bảo trì Partition

### Xóa dữ liệu cũ
```sql
-- Cách 1: DROP luôn (nhanh, không thể phục hồi)
DROP TABLE measurement_y2006m02;

-- Cách 2: Tách ra (vẫn giữ dữ liệu để backup)
ALTER TABLE measurement DETACH PARTITION measurement_y2006m02;
ALTER TABLE measurement DETACH PARTITION measurement_y2006m02 CONCURRENTLY; -- ít lock hơn
```

### Thêm partition mới
```sql
-- Cách 1: Tạo trực tiếp
CREATE TABLE measurement_y2008m02 PARTITION OF measurement
    FOR VALUES FROM ('2008-02-01') TO ('2008-03-01');

-- Cách 2: Tạo bảng riêng → load data → rồi ATTACH (ít lock hơn, an toàn hơn)
CREATE TABLE measurement_y2008m02
    (LIKE measurement INCLUDING DEFAULTS INCLUDING CONSTRAINTS);

ALTER TABLE measurement_y2008m02 ADD CONSTRAINT y2008m02
    CHECK (logdate >= DATE '2008-02-01' AND logdate < DATE '2008-03-01');

\copy measurement_y2008m02 from 'measurement_y2008m02'

ALTER TABLE measurement ATTACH PARTITION measurement_y2008m02
    FOR VALUES FROM ('2008-02-01') TO ('2008-03-01');
```
> 💡 Nên thêm CHECK constraint trước khi ATTACH để tránh scan toàn bộ bảng

### Tạo index không downtime
```sql
-- B1: Tạo index "ảo" trên bảng cha (invalid)
CREATE INDEX measurement_usls_idx ON ONLY measurement (unitsales);

-- B2: Tạo index thật trên từng partition (CONCURRENTLY = không lock)
CREATE INDEX CONCURRENTLY measurement_usls_200602_idx
    ON measurement_y2006m02 (unitsales);

-- B3: Gắn vào parent index → tự động valid khi tất cả partition xong
ALTER INDEX measurement_usls_idx
    ATTACH PARTITION measurement_usls_200602_idx;
```

---

## 5. Partition Pruning (Cắt tỉa phân vùng)

Cơ chế **tự động bỏ qua** các partition không liên quan khi thực thi query.

```sql
-- Bật (mặc định)
SET enable_partition_pruning = on;

-- Kiểm tra query plan
EXPLAIN SELECT count(*) FROM measurement WHERE logdate >= DATE '2008-01-01';
```

- **Pruning xảy ra lúc:** planning time + execution time
- **Không cần index** để pruning hoạt động — chỉ cần partition key trong WHERE clause
- Có thể xem số partition bị bỏ qua qua `"Subplans Removed"` trong EXPLAIN output

---

## 6. Inheritance Partitioning (Cách cũ)

Linh hoạt hơn nhưng phức tạp và chậm hơn. Vẫn còn dùng khi cần:
- Bảng con có thêm cột riêng
- Đa kế thừa
- Logic phân vùng tùy biến

```sql
-- Bảng gốc
CREATE TABLE measurement ( city_id int not null, logdate date not null, ... );

-- Bảng con kế thừa + CHECK constraint
CREATE TABLE measurement_y2006m02 (
    CHECK ( logdate >= DATE '2006-02-01' AND logdate < DATE '2006-03-01' )
) INHERITS (measurement);

-- Cần tạo trigger để route INSERT đúng partition
```

> ❌ Với inheritance: cần viết trigger/rule thủ công, dễ sai, `ON CONFLICT` không hoạt động đúng

---

## 7. Constraint Exclusion vs Partition Pruning

| | Partition Pruning | Constraint Exclusion |
|--|---|---|
| Dùng cho | Declarative partition | Inheritance partition |
| Khi nào | Planning + Execution | Planning only |
| Dựa vào | Partition bounds | CHECK constraints |
| Cấu hình | `enable_partition_pruning` | `constraint_exclusion = partition` |

---

## 8. Giới hạn cần nhớ

- Không thể convert bảng thường → partitioned table (phải tạo mới)
- UNIQUE / PRIMARY KEY phải **bao gồm tất cả partition key columns**
- BEFORE ROW trigger trên INSERT không thể đổi partition đích
- Không mix temporary và permanent table trong cùng partition tree
- `TRUNCATE ONLY` trên bảng cha sẽ báo lỗi

---

## 9. Best Practices

✅ **Chọn cột partition** → cột hay xuất hiện nhất trong WHERE clause  
✅ **Số lượng partition** → vài trăm là ổn, tránh hàng nghìn (tăng planning time + memory)  
✅ **Xóa dữ liệu cũ** → `DETACH` rồi backup trước khi `DROP`  
✅ **Thêm partition mới** → tạo riêng + `ATTACH` để giảm lock thay vì `CREATE ... PARTITION OF`  
✅ **Data warehouse** → có thể dùng nhiều partition hơn OLTP vì planning time ít quan trọng hơn  
✅ **Nên lên kế hoạch từ đầu** → re-partition dữ liệu lớn sau này rất tốn thời gian  

> 💡 **Nhớ:** Nhiều partition hơn không có nghĩa là tốt hơn — cần đánh giá theo workload thực tế