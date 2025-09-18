<script setup lang="ts">
import { MyContainer } from '@jinyaoma/my-components'
defineOptions({
  name: 'PageContainer'
})
</script>

<template>
  <div>
    <MyContainer width="800px">
      <MyContainer.Header height="60px" style="background-color: aquamarine">
        <div>123</div>
      </MyContainer.Header>
      <MyContainer>
        <MyContainer.Aside width="280px" style="background-color: burlywood">
          <div>456</div>
        </MyContainer.Aside>
        <MyContainer height="400px" scrollable>
          <MyContainer.Main style="min-height: 1000px; background-color: greenyellow">
            <div>789</div>
          </MyContainer.Main>
          <MyContainer.Footer style="background-color: chartreuse">
            <div>abc</div>
          </MyContainer.Footer>
        </MyContainer>
      </MyContainer>
    </MyContainer>
  </div>
</template>

<style scoped lang="scss"></style>
