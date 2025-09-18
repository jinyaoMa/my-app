export enum IconSize {
  regular,
  small,
  large
}
export type IconSizeKey = keyof typeof IconSize
export const IconSizeHelper = {
  isEnum: (s: IconSize | IconSizeKey): s is IconSize => typeof s !== 'string' && s in IconSize,
  toEnum: (s: IconSize | IconSizeKey) => (IconSizeHelper.isEnum(s) ? s : IconSize[s]),
  toString: (s: IconSize | IconSizeKey) => (IconSizeHelper.isEnum(s) ? IconSize[s] : s)
}
