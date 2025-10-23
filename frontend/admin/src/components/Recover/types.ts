export interface RecoverEntityId {
  id?: string | undefined
}

export type RecoverEntityIdReserved = RecoverEntityId & {
  reserved?: boolean
}
