export enum Mode {
  normal,
  medium,
  mini
}
export type ModeKey = keyof typeof Mode
export const ModeHelper = {
  isEnum: (m: Mode | ModeKey): m is Mode => typeof m !== 'string' && m in Mode,
  toEnum: (m: Mode | ModeKey) => (ModeHelper.isEnum(m) ? m : Mode[m]),
  toString: (m: Mode | ModeKey) => (ModeHelper.isEnum(m) ? Mode[m] : m)
}

export enum ColorType {
  primary,
  success,
  warning,
  error,
  info,
  link
}
export type ColorTypeKey = keyof typeof ColorType
export const ColorTypeHelper = {
  isEnum: (t: ColorType | ColorTypeKey): t is ColorType => typeof t !== 'string' && t in ColorType,
  toEnum: (t: ColorType | ColorTypeKey) => (ColorTypeHelper.isEnum(t) ? t : ColorType[t]),
  toString: (t: ColorType | ColorTypeKey) => (ColorTypeHelper.isEnum(t) ? ColorType[t] : t)
}
