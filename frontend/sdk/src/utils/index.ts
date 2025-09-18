import { request } from './request'
import { requestHeaders, responseError, responseHook } from './interceptors'

export * from './request'

request.useRequestInterceptor(requestHeaders)
request.useResponseInterceptor(responseHook, responseError)
