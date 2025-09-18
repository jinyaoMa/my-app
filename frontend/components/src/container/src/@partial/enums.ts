export enum ContainerDirection {
  vertical,
  horizontal
}
export type ContainerDirectionKey = keyof typeof ContainerDirection
export const ContainerDirectionHelper = {
  isEnum: (d: ContainerDirection | ContainerDirectionKey): d is ContainerDirection =>
    typeof d !== 'string' && d in ContainerDirection,
  toEnum: (d: ContainerDirection | ContainerDirectionKey) =>
    ContainerDirectionHelper.isEnum(d) ? d : ContainerDirection[d],
  toString: (d: ContainerDirection | ContainerDirectionKey) =>
    ContainerDirectionHelper.isEnum(d) ? ContainerDirection[d] : d
}

export enum ContainerSticky {
  top,
  bottom,
  left,
  right
}
export type ContainerStickyKey = keyof typeof ContainerSticky
export const ContainerStickyHelper = {
  isEnum: (s: ContainerSticky | ContainerStickyKey): s is ContainerSticky =>
    typeof s !== 'string' && s in ContainerSticky,
  toEnum: (s: ContainerSticky | ContainerStickyKey) =>
    ContainerStickyHelper.isEnum(s) ? s : ContainerSticky[s],
  toString: (s: ContainerSticky | ContainerStickyKey) =>
    ContainerStickyHelper.isEnum(s) ? ContainerSticky[s] : s,
  toUniqueStringArray: (
    s?: ContainerSticky | ContainerStickyKey | (ContainerSticky | ContainerStickyKey)[]
  ) => [
    ...new Set(
      Array.isArray(s)
        ? s.map(ContainerStickyHelper.toString)
        : s
        ? [ContainerStickyHelper.toString(s)]
        : []
    )
  ]
}
