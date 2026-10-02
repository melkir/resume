import adapter from '@sveltejs/adapter-static';
import { relative, sep } from 'node:path';
import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

const dev = process.argv.includes('dev');

export default defineConfig({
  plugins: [
    tailwindcss(),
    sveltekit({
      compilerOptions: {
        runes: ({ filename }) => {
          const relativePath = relative(import.meta.dirname, filename);
          const pathSegments = relativePath.toLowerCase().split(sep);
          const isExternalLibrary = pathSegments.includes('node_modules');

          return isExternalLibrary ? undefined : true;
        }
      },
      adapter: adapter(),
      paths: { base: dev ? '' : process.env.BASE_PATH }
    })
  ]
});
