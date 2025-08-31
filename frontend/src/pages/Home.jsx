import React from "react";

const ebooks = [
  {
    title: "The Art of Go",
    author: "Jane Doe",
    cover: "https://images.unsplash.com/photo-1512820790803-83ca734da794?auto=format&fit=crop&w=200&q=80",
    description: "A comprehensive guide to mastering the Go programming language.",
  },
  {
    title: "React in Action",
    author: "John Smith",
    cover: "https://images.unsplash.com/photo-1524985069026-dd778a71c7b4?auto=format&fit=crop&w=200&q=80",
    description: "Learn React from scratch with hands-on projects and best practices.",
  },
  {
    title: "Fullstack Fundamentals",
    author: "Alex Johnson",
    cover: "https://images.unsplash.com/photo-1465101046530-73398c7f28ca?auto=format&fit=crop&w=200&q=80",
    description: "Everything you need to know to build modern web applications.",
  },
];

function Home() {
  return (
    <div style={{ padding: "2rem", background: "#f7f7f7", minHeight: "100vh" }}>
      <h1 style={{ textAlign: "center", marginBottom: "2rem" }}>📚 FableForge eBooks</h1>
      <div style={{
        display: "flex",
        flexWrap: "wrap",
        gap: "2rem",
        justifyContent: "center"
      }}>
        {ebooks.map((book, idx) => (
          <div key={idx} style={{
            background: "#fff",
            borderRadius: "12px",
            boxShadow: "0 2px 8px rgba(0,0,0,0.1)",
            width: "250px",
            padding: "1rem",
            display: "flex",
            flexDirection: "column",
            alignItems: "center"
          }}>
            <img
              src={book.cover}
              alt={book.title}
              style={{ width: "120px", height: "160px", objectFit: "cover", borderRadius: "8px", marginBottom: "1rem" }}
            />
            <h2 style={{ fontSize: "1.2rem", margin: "0.5rem 0" }}>{book.title}</h2>
            <p style={{ color: "#555", margin: "0.25rem 0 0.5rem 0" }}>by {book.author}</p>
            <p style={{ fontSize: "0.95rem", color: "#777", textAlign: "center" }}>{book.description}</p>
            <button style={{
              marginTop: "1rem",
              padding: "0.5rem 1.5rem",
              background: "#1976d2",
              color: "#fff",
              border: "none",
              borderRadius: "6px",
              cursor: "pointer"
            }}>
              Read Now
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Home;