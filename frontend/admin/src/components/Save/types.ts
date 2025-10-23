export interface SaveEntityId {
  id?: string | undefined
}

export type SaveEntityIdReserved = SaveEntityId & {
  reserved?: boolean
}

export type SaveDetailRequestData<T> = {
  data: T | undefined
  success?: boolean | undefined
  total?: number | undefined
} & Record<string, any>
