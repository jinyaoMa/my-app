export interface QueryEntityId {
  id?: string | undefined
}

export type QueryEntityIdReserved = QueryEntityId & {
  reserved?: boolean
}
