db = db.getSiblingDB("grpcdb"); // Tạo hoặc chọn cơ sở dữ liệu 'mydatabase'
db.createCollection("posts"); // Tạo collection 'posts'

db.posts.insertMany([ // Chèn dữ liệu mẫu vào collection 'posts'
    {
        Title: "First Post",
        Content: "This is the first post",
        Image: "image1.jpg",
        User: "user1",
        CreateAt: new Date(),
        UpdateAt: new Date()
      }
]);