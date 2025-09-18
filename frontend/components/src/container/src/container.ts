import { Mode } from '../../@global'
import type { ModeKey } from '../../@global'
import { ContainerDirection } from './@partial/enums'
import type { ContainerDirectionKey } from './@partial/enums'

export interface ContainerProps {
  /**
   * @description display mode
   */
  mode?: Mode | ModeKey
  /**
   * @description container direction
   */
  direction?: ContainerDirection | ContainerDirectionKey
  /**
   * @description is container viewport of browser
   */
  viewport?: boolean
  /**
   * @description container height
   */
  height?: string | number
  /**
   * @description container width
   */
  width?: string | number
  /**
   * @description is container scrollable
   */
  scrollable?: boolean
}
