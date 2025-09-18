import { ContainerSticky } from './@partial/enums'
import type { ContainerStickyKey } from './@partial/enums'

export interface HeaderProps {
  /**
   * @description height of header
   */
  height?: string | number
  /**
   * @description is header sticky, and header sticky type
   */
  sticky?: ContainerSticky | ContainerStickyKey | (ContainerSticky | ContainerStickyKey)[]
}
