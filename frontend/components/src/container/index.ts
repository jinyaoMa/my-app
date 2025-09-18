import { withInstall, withNoopInstall } from '../withInstall'
import Container from './src/container.vue'
import Aside from './src/aside.vue'
import Header from './src/header.vue'
import Footer from './src/footer.vue'
import Main from './src/main.vue'

export const MyAside = withNoopInstall(Aside)
export const MyHeader = withNoopInstall(Header)
export const MyFooter = withNoopInstall(Footer)
export const MyMain = withNoopInstall(Main)
export const MyContainer = withInstall(Container, {
  Aside: MyAside,
  Header: MyHeader,
  Footer: MyFooter,
  Main: MyMain
})

export default MyContainer
