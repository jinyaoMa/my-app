export interface DetailEntityId {
  id?: string | undefined
}

export type DetailEntityIdReserved = DetailEntityId & {
  reserved?: boolean
}

export type DetailRequestData<T> = {
  data: T | undefined
  success?: boolean | undefined
  total?: number | undefined
} & Record<string, any>
