import { sveltekit } from '@sveltejs/kit/vite';
import path from 'path';

/** @type {import('vite').UserConfig} */
const config = {
    plugins: [sveltekit()],
    server: {
        port: 3000,
        host: true,
    },
    resolve: {
        alias: {
            src: path.resolve(__dirname, 'src'),
        },
    }
};

export default config;
