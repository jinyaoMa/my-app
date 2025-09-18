import { generateService } from '@umijs/openapi'

generateService({
  requestLibPath: "import { request } from '../utils'",
  schemaPath: 'https://localhost:18443/docs/openapi.json',
  serversPath: './src'
})
