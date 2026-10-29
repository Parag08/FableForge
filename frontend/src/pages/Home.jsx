import React from "react";

// PLACEHOLDER: static mock of the planned pipeline. Nothing here is wired to
// the backend yet; see the roadmap in the repository README.
const pipelineSteps = [
  {
    title: "1. Upload a chapter",
    description: "Paste or upload a chapter of your novel as plain text.",
  },
  {
    title: "2. Scenes + character bible",
    description:
      "The chapter is split into illustratable scenes, and recurring characters are described once in a shared character bible.",
  },
  {
    title: "3. Consistent illustrations",
    description:
      "Each scene is illustrated using the character bible so faces, outfits and settings stay consistent across images.",
  },
  {
    title: "4. Export EPUB",
    description: "Text and illustrations are assembled into an illustrated EPUB.",
  },
];

function Home() {
  return (
    <div style={{ padding: "2rem", background: "#f7f7f7", minHeight: "100vh" }}>
      <h1 style={{ textAlign: "center", marginBottom: "0.5rem" }}>FableForge</h1>
      <p style={{ textAlign: "center", color: "#555", marginBottom: "0.5rem" }}>
        Turn a novel into a consistently illustrated book.
      </p>
      <p
        style={{
          textAlign: "center",
          color: "#8a5a00",
          background: "#fff4d6",
          borderRadius: "6px",
          padding: "0.5rem",
          maxWidth: "560px",
          margin: "0 auto 2rem auto",
        }}
      >
        Placeholder: this page previews the planned workflow. None of these steps are implemented yet.
      </p>
      <div style={{ display: "flex", flexWrap: "wrap", gap: "1.5rem", justifyContent: "center" }}>
        {pipelineSteps.map((step) => (
          <div
            key={step.title}
            style={{
              background: "#fff",
              borderRadius: "12px",
              boxShadow: "0 2px 8px rgba(0,0,0,0.1)",
              width: "240px",
              padding: "1rem",
            }}
          >
            <h2 style={{ fontSize: "1.1rem", margin: "0.25rem 0 0.75rem 0" }}>{step.title}</h2>
            <p style={{ fontSize: "0.95rem", color: "#666", margin: 0 }}>{step.description}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Home;
