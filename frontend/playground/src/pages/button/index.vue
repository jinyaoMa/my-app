<script setup lang="ts">
import { MyContainer, MyButton } from '@jinyaoma/my-components'
defineOptions({
  name: 'PageButton'
})
</script>

<template>
  <MyContainer
    :style="{
      lineHeight: '2em'
    }"
  >
    <MyContainer.Main>
      测试 <MyButton icon-name="jinyao-ma" round>BUTTON</MyButton><br />
      123
    </MyContainer.Main>
  </MyContainer>
</template>
