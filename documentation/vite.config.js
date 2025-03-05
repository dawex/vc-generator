import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

import { viteStaticCopy } from 'vite-plugin-static-copy'

function getBaseUrl(env) {
  if (env === 'production') {
    return "/vc-generator/"
  }
	return "/";
}

// https://vitejs.dev/config/
export default defineConfig({
  base: getBaseUrl(process.env.NODE_ENV),
  define: {
    '__APP_BASE_URL__': `"${getBaseUrl(process.env.NODE_ENV)}"`,
  },
  server: {
    host: '127.0.0.1'
  },
  plugins: [
    vue(),
    viteStaticCopy({
      targets: [
        {
          src: `oas/dist/*.yaml`,
          dest: ''
        },
        {
          src: `redoc_config/.redocly.yaml`,
          dest: ''
        }
      ]
    })
  ],
  resolve: {
    alias: [
      {
        find: /^~.+/,
        replacement: (val) => {
          return val.replace(/^~/, "");
        },
      },
    ],
  },
})
