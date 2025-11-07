1. Cơ bản về cú pháp và kiểu dữ liệu

👉 Mục tiêu: hiểu cú pháp, cách khai báo, nhập/xuất và điều kiện.

Bài tập:

[x] In ra “Hello, World”

[x] Nhập tên người dùng rồi in “Hello, {name}”

[x] Tính tổng, hiệu, tích, thương của hai số nhập từ bàn phím

 Viết hàm kiểm tra số chẵn/lẻ

 Viết hàm kiểm tra năm nhuận

🔁 2. Vòng lặp & điều kiện

👉 Làm quen với for, if, switch.

Bài tập:

 In bảng cửu chương 1–9

 Tính tổng các số từ 1 đến n

 Tìm số lớn nhất trong một mảng

 Đếm số lần xuất hiện của một ký tự trong chuỗi

 Viết chương trình kiểm tra số nguyên tố

🧱 3. Mảng (Array), Slice và Map

👉 Làm quen với cấu trúc dữ liệu phổ biến của Go.

Bài tập:

 Viết hàm đảo ngược một slice

 Tính tổng các phần tử trong slice

 Dùng map để đếm tần suất từ trong một câu

 Viết chương trình quản lý danh bạ (map[string]string)

 Lọc các phần tử trùng trong slice

🧰 4. Hàm (Function) và Struct

👉 Làm quen với lập trình có cấu trúc.

Bài tập:

 Viết hàm nhận 2 số và trả về giá trị lớn nhất

 Viết hàm trả về nhiều giá trị (VD: tổng và trung bình)

 Định nghĩa struct Person có Name và Age, in thông tin người

 Viết method SayHello() cho struct Person

 Tạo struct Student kế thừa Person (embed struct)

📦 5. Package & Module

👉 Học cách tổ chức code gọn gàng.

Bài tập:

 Tách logic ra package utils và import vào main.go

 Viết package mathhelper chứa hàm Sum, Average, Min, Max

 Dùng go mod init, go get để quản lý module

⚙️ 6. Error Handling

👉 Biết cách xử lý lỗi đúng kiểu Go.

Bài tập:

 Viết hàm chia 2 số và trả lỗi nếu chia cho 0

 Dùng errors.New() hoặc fmt.Errorf() để tạo lỗi

 Thử defer, panic, recover để hiểu luồng lỗi

🧵 7. Goroutine & Channel (Concurrency)

👉 Là phần rất đặc trưng của Go.

Bài tập:

 Chạy 2 hàm song song in số chẵn/lẻ

 Dùng channel để trao đổi dữ liệu giữa 2 goroutine

 Tạo worker pool đơn giản (5 goroutine cùng xử lý job)

 Sử dụng sync.WaitGroup để đợi tất cả goroutine hoàn thành

🌐 8. File I/O & JSON

👉 Làm quen thao tác với dữ liệu.

Bài tập:

 Ghi và đọc nội dung từ file .txt

 Lưu slice thành JSON file

 Parse JSON vào struct

 Đọc dữ liệu từ file cấu hình (config.json)

🧠 9. REST API cơ bản

👉 Áp dụng Go vào web development.

Bài tập:

 Tạo API GET /hello trả về chuỗi JSON

 Tạo API CRUD quản lý User (thêm, sửa, xóa, xem)

 Dùng thư viện net/http hoặc framework gin-gonic/gin

 Trả lỗi HTTP hợp lý (404, 500, v.v.)

🧪 10. Unit Test & Benchmark

👉 Rất quan trọng khi làm dự án thực tế.

Bài tập:

 Viết file math_test.go test hàm Sum()

 Dùng t.Errorf() để xác nhận kết quả sai

 Viết benchmark test (go test -bench=.)

💼 11. Bài tập nâng cao (ứng dụng thực tế)

👉 Khi Ellie đã nắm vững nền tảng.

Gợi ý dự án nhỏ:

✅ Todo App CLI (command line)

✅ REST API quản lý bài viết (Post CRUD)

✅ Chat server đơn giản với Goroutine

✅ Worker xử lý hàng đợi (queue processing)

✅ Go microservice nói chuyện qua gRPC