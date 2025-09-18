import { ColorType } from '../../@global'
import type { ColorTypeKey } from '../../@global'

export interface ButtonProps {
  /**
   * @description display mode
   */
  text?: string
  /**
   * @description icon name
   */
  iconName?: string
  /**
   * @description round button
   */
  round?: boolean
  /**
   * @description button color type
   */
  type?: ColorType | ColorTypeKey
}
