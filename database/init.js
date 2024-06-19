db = db.getSiblingDB("grpcdb"); // Tạo hoặc chọn cơ sở dữ liệu 'mydatabase'
db.createCollection("posts"); // Tạo collection 'posts'
db.createCollection("users");
db.posts.insertMany([ // Chèn dữ liệu mẫu vào collection 'posts'
  {
    title: "First Post",
    content: "This is the first post",
    image: "image1.jpg",
    user: "user1",
    create_at: new Date(),
    updated_at: new Date()
  },
  {
    title: "First Post",
    content: "This is the first post",
    image: "image1.jpg",
    user: "user1",
    create_at: new Date(),
    updated_at: new Date()
  },
  {
    title: "First Post",
    content: "This is the first post",
    image: "image1.jpg",
    user: "user1",
    create_at: new Date(),
    updated_at: new Date()
  },
  {
    title: "First Post",
    content: "This is the first post",
    image: "image1.jpg",
    user: "user1",
    create_at: new Date(),
    updated_at: new Date()
  },
  {
    title: "First Post",
    content: "This is the first post",
    image: "image1.jpg",
    user: "user1",
    create_at: new Date(),
    updated_at: new Date()
  },
  {
    title: "First Post",
    content: "This is the first post",
    image: "image1.jpg",
    user: "user1",
    create_at: new Date(),
    updated_at: new Date()
  },
  {
    title: "First Post",
    content: "This is the first post",
    image: "image1.jpg",
    user: "user1",
    create_at: new Date(),
    updated_at: new Date()
  }
]);
db.users.insertMany([
  {
    username: "user1",
    email: "email1",
    password: "password1",
  },
  {
    username: "user2",
    email: "email1",
    password: "password1",
  },
  {
    username: "user3",
    email: "email1",
    password: "password1",
  },
  {
    username: "user4",
    email: "email1",
    password: "password1",
  }

]);