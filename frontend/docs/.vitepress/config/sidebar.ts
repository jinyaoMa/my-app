import { createComponentSiderBar } from "../utils/createComponentSiderBar"

const _componentsSiderBarItems:{
  text: string
  link: string
}[] = [

]

const componentsSiderBarItems = createComponentSiderBar() ?? _componentsSiderBarItems

export const sidebar = {
  '/guide/': [
    {
      text: '指引',
      items: [
        {
          text: '特性',
          link: '/guide/features'
        },
        {
          text: '快速上手',
          link: '/guide/started'
        },
        {
          text: '开发攻略',
          link: '/guide/go'
        }
      ]
    },
    {
      text: '组件',
      items: componentsSiderBarItems,
    }
  ],
}