import { ContainerSticky } from './@partial/enums'
import type { ContainerStickyKey } from './@partial/enums'

export interface FooterProps {
  /**
   * @description height of footer
   */
  height?: string | number
  /**
   * @description is footer sticky, and footer sticky type
   */
  sticky?: ContainerSticky | ContainerStickyKey | (ContainerSticky | ContainerStickyKey)[]
}
