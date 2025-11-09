import { defineConfig } from '@kubb/core'

export default defineConfig(() => {
  return {
    root: '.',
    input: {
      path: '../server/src/docs/swagger.yaml',
    },
    output: {
      path: './utils/backend_types',
    },
    plugins: [],
  }
})