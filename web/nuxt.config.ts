// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },

  features: {
    devLogs: true
  },

  experimental: {
    entryImportMap: false
  },

  modules: [
    '@pinia/nuxt',
  ],

  runtimeConfig: {
    public: {
      apiBase: 'http://localhost:8080',
    },
  },

  typescript: {
    tsConfig: {
      compilerOptions: {
        noEmit: true,
        allowImportingTsExtensions: true,
      }
    },
    nodeTsConfig: {
      include: ['../kubb.config.ts']
    }
  },

  modules: ['@pinia/nuxt']
})