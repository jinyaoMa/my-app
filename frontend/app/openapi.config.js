import { generateService } from '@umijs/openapi'

generateService({
  requestLibPath: "import request from '@/utils/request'",
  schemaPath: 'https://localhost:18443/apid/openapi.json',
  serversPath: './src/services'
})
