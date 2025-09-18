import { ContainerSticky } from './@partial/enums'
import type { ContainerStickyKey } from './@partial/enums'

export interface AsideProps {
  /**
   * @description width of aside
   */
  width?: string | number
  /**
   * @description is aside sticky, and aside sticky type
   */
  sticky?: ContainerSticky | ContainerStickyKey | (ContainerSticky | ContainerStickyKey)[]
}
