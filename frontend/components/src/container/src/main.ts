import { ContainerSticky } from './@partial/enums'
import type { ContainerStickyKey } from './@partial/enums'

export interface MainProps {
  /**
   * @description is main sticky, and main sticky type
   */
  sticky?: ContainerSticky | ContainerStickyKey | (ContainerSticky | ContainerStickyKey)[]
}
