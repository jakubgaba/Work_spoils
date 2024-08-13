
import adapter from '@sveltejs/adapter-auto';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
// Remove the following import if not used:
// import { preprocess } from 'svelte/compiler';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Preprocess should be at the root level of the config object
	preprocess: vitePreprocess({
		scss: {
			prependData: `@import './src/styles.scss';`
		},
	}),
	kit: {
		// Corrected the typo in the comment
		// adapter-auto only supports some environments, see https://kit.svelte.dev/docs/adapter-auto for a list.
		// If your environment is not supported, or you settled on a specific environment, switch out the adapter.
		// See https://kit.svelte.dev/docs/adapters for more information about adapters.
		adapter: adapter({
			pages: 'build',
			assets: 'build',
			fallback: 'index.html',
			precompress: true,
			strict: true,
		}),
	},
};

export default config;
