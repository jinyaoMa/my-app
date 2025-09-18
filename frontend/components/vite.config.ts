import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import dts from 'vite-plugin-dts'

export default defineConfig({
  build: {
    target: 'modules',
    outDir: 'es',
    emptyOutDir: true,
    minify: true,
    lib: {
      entry: 'src/index.ts',
      name: 'MyDesign',
      fileName: (format) => `index.${format}.js`
    },
    rollupOptions: {
      external: ['vue'],
      output: [
        {
          exports: 'named',
          format: 'es',
          dir: '../design/dist/es',
          entryFileNames: '[name].js',
          preserveModules: true,
          preserveModulesRoot: 'src'
        },
        {
          exports: 'named',
          format: 'cjs',
          dir: '../design/dist/lib',
          entryFileNames: '[name].js',
          preserveModules: true,
          preserveModulesRoot: 'src'
        }
        // {
        //   exports: 'named',
        //   format: 'umd',
        //   dir: '../design/dist/umd',
        //   entryFileNames: '[name].js',
        //   name: 'MyDesign',
        //   globals: {
        //     vue: 'Vue'
        //   }
        // }
      ]
    }
  },
  plugins: [
    vue(),
    dts({
      tsconfigPath: 'tsconfig.json',
      entryRoot: './src',
      exclude: ['node_modules', 'vite.config.ts'],
      outDir: '../design/dist/es'
    }),
    dts({
      tsconfigPath: 'tsconfig.json',
      entryRoot: './src',
      exclude: ['node_modules', 'vite.config.ts'],
      outDir: '../design/dist/lib'
    })
  ]
})
