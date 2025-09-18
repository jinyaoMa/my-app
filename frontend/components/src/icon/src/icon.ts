import { IconSize } from './@partial/enums'
import type { IconSizeKey } from './@partial/enums'

export interface IconProps {
  /**
   * @description icon name
   */
  name: string
  /**
   * @description icon size
   */
  size?: IconSize | IconSizeKey
  /**
   * @description fix width
   */
  fixWidth?: string
}
