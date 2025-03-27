# VCS-Redis

## Overview
Redis (Remote dictionary server) là 1 kho lưu trữ dữ liệu mã nguồn mở chủ yếu được sử dụng như 1 cơ sở dữ liệu, cache và message queue. Redis được biết đến với khả năng sẵn sàng cao, độ trễ thấp và tính linh hoạt trong việc xử lý với nhiều kiểu dữ liệu khác nhau. Redis lưu trữ dữ liệu trên RAM thay vì trên ổ đĩa giúp việc xử lý dữ liệu nhanh hơn rất nhiều so với các cơ sở dữ liệu truyền thống khác.

Redis được sử dụng cho các mục đích sau:
* **Cache**: Tăng tốc ứng dụng bằng cách cache các dữ liệu xuất hiện thường xuyên.
* **Session Management**: Lưu trữ phiên của người dùng trên ứng dụng web.
* **Realtime Analytics**: Sử dụng trong bảng xếp hạng hoặc theo dõi số liệu
* **Pub-sub Messaging**: Hỗ trợ kiến trúc event driving.

## Caching
Redis hỗ trợ nhiều chiến lược cache để kiểm soát dữ liệu hiệu quả, tránh việc sử dụng bộ nhớ quá mức, và vẫn giữ lại được những dữ liệu được truy cập thường xuyên.
### Cache-Aside
Chiến lược này còn được gọi là lazy loading. Vùng nhớ cache (Redis) được chạy song song với database gốc, do đó ứng dụng có thể tương tác trực tiếp với cache và database.

#### Cách thức hoạt động:
1. Application kiểm tra trong Redis xem có lưu trữ dữ liệu mình cần không.
2. Khi cache không chứa dữ liệu mà application cần (cache miss), pplication sẽ xuống database để lấy dữ liệu.
3. Application sẽ lưu dữ liệu lấy được từ database để lưu vào cache.

#### Lợi ích:
1. Đảm bảo cache được các dữ liệu được truy cập thường xuyên.
2. Tránh việc lưu trữ các dữ liệu không cần thiết.
3. Ứng dụng có thể truy cập thoải mái vào cache và database. Nếu cache không hoạt động thì ứng dụng vẫn chạy bình thường.

#### Bất lợi:
1. Dữ liệu được truy cập lần đầu hoặc khi hết hạn tốn nhiều thời gian hơn.
2. Khi dữ liệu trong database thay đổi thì cache có thể không còn đúng.

### Read through
Thay vì ứng dụng phải kết nối với cả cache và database, **Read Through** cho phép ứng dụng chỉ cần truy cập vào cache, còn lại là do cache xử lý.

#### Cách thức hoạt động:
1. Ứng dụng gửi request đến cache để lấy dữ liệu
2. Nếu cache hit, cache sẽ gửi dữ liệu ngay lập tức cho ứng dụng.
3. Nếu cache miss, cache sẽ truy cập vào database để lấy dữ liệu, sau đó cập nhật dữ liệu của chính nó và gửi dữ liệu về cho ứng dụng.

#### Lợi ích
1. Ứng dụng không cần quan tâm đến việc cache dữ liệu, mọi thứ đều do cache xử lý.
2. Đảm bảo cache được các dữ liệu truy cập thường xuyên.

#### Bất lợi
1. Phụ thuộc nhiều vào cache, nếu cache không hoạt động thì dữ liệu cũng sẽ sập theo.
2. Không có quyền lựa chọn thời gian cache dữ liệu.
3. Khi dữ liệu trong database thay đổi thì cache có thể không còn đúng.

### Write-through cache
Khi ứng dụng cần ghi dữ liệu, thay vì ghi dữ liệu vào database, nó sẽ ghi vào cache trước và cache sẽ ghi vào database sau. Đây là cách ghi dữ liệu **đồng bộ (synchronously)**.

#### Cách thức hoạt động
Khi một request write tới:

1. Dữ liệu sẽ được lưu vào cache
2. Cache sẽ gửi yêu cầu lưu dữ liệu vào database ngay lập tức.

#### Lợi ích
1. Không xảy ra cache miss do dữ liệu luôn được ghi vào cache trước khi vào database.
2. Đồng nhất được cache và database.
3. Kết hợp được với read through cache.

#### Bất lợi
1. Hầu hết dữ liệu write đều chỉ dùng 1 lần nên dễ ghi các dữ liệu không cần thiết.
2. Chỉ thích hợp với write-heavy workloads.

### Write back
Thay vì cache ghi dữ liệu trực tiếp vào database khi nhận được request, cache sẽ đồng bộ dữ liệu xuống database định kì theo thời gian, hoặc theo số lượng dữ liệu được insert/update. Đây là cách ghi dữ liệu **bất đồng bộ (asynchronously)**.

#### Cách thức hoạt động
1. Dữ liệu sẽ được lưu vào cache
2. Sau một khoảng thời gian, cache sẽ ghi dữ liệu vào database.

#### Lợi ích
1. Giảm tải áp lực write xuống database, từ đó sẽ giảm được chi phí và các vấn đề liên quan tới database
2. Kết hợp được với read through cache.

#### Bất lợi
1. Eventual consistency, database có một khoảng thời gian không được đồng bộ với cache.
2. Nếu cache sập thì ứng dụng cũng sập và sẽ bị mất toàn bộ dữ liệu chưa kịp đồng bộ vào database.

### Write around
#### Cách thức hoạt động
1. Dữ liệu chỉ được ghi vào database.
2. Khi ứng dụng muốn đọc dữ liệu, đầu tiên sẽ đọc ở cache. Nếu cache hit thì trả dữ liệu về.
3. Nếu cache miss, ứng dụng sẽ đọc dữ liệu từ database.
4. Ứng dụng ghi dữ liệu vào cache.

#### Lợi ích
1. Tránh ghi các dữ liệu không cần thiết.
2. Phù hợp với dữ liệu không cần truy cập nhiều nhưng vẫn cần cache trong lúc đọc.

#### Bất lợi
1. Không xử lý được với dữ liệu được update nhiều.

## Data Structures
### Strings (SET, GET)
Là kiểu dữ liệu cơ bản nhất trong Redis, lưu dữ liệu dưới dạng key-value. Thường dùng để cache các thông tin xuất hiện thường xuyên.

Ví dụ:
```redis
SET "user 123" "John Doe"
GET "user 123"					# John Doe
```

### Lists (LPUSH, RPUSH, LPOP, RPOP, LRANGE)
Lưu trữ một tập hợp string, được cài đặt như một linked list. Thường được sử dụng như message queue hay hàng đợi cho job hoặc công việc.

Ví dụ:
```redis
LPUSH tasks "task1" "task2" "task3"   # Push tasks to the left (like a stack)
LRANGE tasks 0 -1                     # Get all tasks
RPOP tasks                             # Pop a task from the right
```

### Sets (SADD, SREM, SMEMBERS)
Tập hợp các giá trị phân biệt.

Ví dụ:
```redis
SADD unique_users "user1" "user2" "user3" # Add 3 users
SMEMBERS unique_users"               # List all users
SREM unique_users "user2"              # Delete user2
SISMEMBER unique_users "user2"  # Check if user2 exists
```

### Sorted Sets (ZADD, ZRANGE, ZREM)
Giống Sets nhưng mỗi giá trị được gán điểm số, sort các giá trị theo điểm số đó. Thường dùng cho bảng rank hoặc sắp xếp các job theo thứ tự ưu tiên.

Ví dụ:
```redis
ZADD leaderboard 100 "player1"
ZADD leaderboard 200 "player2"
ZRANGE leaderboard 0 -1 WITHSCORES  # Get sorted players
```

### Hashes (HSET, HGET, HGETALL)
Một key-value được lưu trữ trong 1 key (giống JSON). Thường sử dụng để lưu trữ dữ liệu người dùng hoặc cache.

Ví dụ:
```redis
HSET user:100 name "Alice" age 30 city "NY"
HGET user:100 name  # Get specific field
HGETALL user:100    # Get all fields
```

### Bitmaps (SETBIT, GETBIT, BITCOUNT)
Biểu diễn dữ liệu dưới dạng nhị phân.

Ví dụ:
```redis
SETBIT user_activity 1 1  # Mark user 1 as active
SETBIT user_activity 2 0  # Mark user 2 as inactive
BITCOUNT user_activity    # Count active users
```

### HyperLogLog (PFADD, PFCOUNT)
Cấu trúc dữ liệu xác suất để đếm các phần tử duy nhất một cách hiệu quả.

Ví dụ:
```redis
PFADD unique_visitors "user1" "user2" "user3"
PFCOUNT unique_visitors  # Approximate count of unique users
```

### Streams (XADD, XREAD, XGROUP)