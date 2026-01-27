export default function HomePage() {
  return (
    <main
      style={{
        maxWidth: "800px",
        margin: "2rem auto",
        padding: "0 1rem",
        fontFamily: "system-ui, sans-serif",
      }}
    >
      <header>
        <h1>It Works!!!</h1>
        <p>
            I'm learning how to write React.js and Next.js. This site will be
            my personal website to show my webdev skills to recruiters and to
            talk about during interviews.
        </p>
      </header>

      <section style={{ marginTop: "1.5rem" }}>
        <h2>⚠️ Currently Under Development ⚠️</h2>
        <ul>
          <li>
            <strong>About Me Page:</strong> A reactive/responsive page that shows what jobs I'm looking for and my experience.
          </li>
          <li>
            <strong>Projects Page:</strong> A reactive/responsive page that shows my projects working with React.js and Next.js.
          </li>
          <li>
            <strong>Blog:</strong> A library of my techinical blog posts to show my subject matter expertise in various areas.
          </li>
        </ul>
      </section>
    </main>
  );
}
