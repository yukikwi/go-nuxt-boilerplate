import { defineConfig } from '@kubb/core'
import { pluginOas } from '@kubb/plugin-oas'
import { pluginTs } from '@kubb/plugin-ts'
import { pluginZod } from '@kubb/plugin-zod'
import { pluginClient } from '@kubb/plugin-client'

export default defineConfig(() => {
  return {
    root: '',
    input: {
      path: '../server/src/static/openapi/docs.yaml',
    },
    output: {
      path: './app/clients',
      clean: true,
      barrelType: false
    },
    plugins: [
      pluginOas({
        output: {
          path: './schemas',
        },
      }),
      pluginTs({
        output: {
          path: './types',
        },
      }),
      pluginZod({
        output: {
          path: './zod',
        },
      }),
      pluginClient({
        baseURL: process.env.NUXT_PUBLIC_API_BASE_URL || 'http://localhost:8081',
        output: {
          path: './axios',
          banner: '// @ts-nocheck'
        },
        group: {
          type: 'tag',
          name: ({ group }) => `${group}Service`,
        },
        transformers: {
          name: (name, type) => {
            return `${name}Client`
          },
        },
        operations: true,
        parser: 'zod',
        pathParamsType: "object",
        dataReturnType: 'full',
        importPath: '@/utils/fetch',
      }),
    ],
  }
})