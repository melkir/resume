# Résumé

A print-friendly résumé site built with SvelteKit, Tailwind CSS, and the static adapter.

## Development

Install dependencies and start the local development server:

```bash
bun install
bun run dev
```

## Validation

Run the same checks used by CI before opening a pull request:

```bash
bun run check
bun run lint
bun run build
```

## Deployment

Pushes to `main` are validated, built, and deployed to GitHub Pages. Pull requests run the
same validation and production build without deploying.
