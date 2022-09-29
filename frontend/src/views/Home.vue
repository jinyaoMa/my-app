<script setup lang="ts">
import { useI18n } from "vue-i18n";

const { t } = useI18n();
</script>

<template>
  <my-container class="home">
    <div class="top">
      <a href="https://wails.io/" target="_blank">
        <img
          src="../assets/wails-logo-textless.svg"
          class="logo wails"
          alt="Wails logo"
        />
      </a>
    </div>
    <div class="banner">
      <a href="https://github.com/jinyaoMa" target="_blank">
        <img src="/icon.svg" class="logo avatar" alt="Author logo" />
      </a>
      <div class="divider"></div>
      <a href="https://vitejs.dev" target="_blank">
        <img src="../assets/vite.svg" class="logo" alt="Vite logo" />
      </a>
      <div class="divider"></div>
      <a href="https://vuejs.org/" target="_blank">
        <img src="../assets/vue.svg" class="logo vue" alt="Vue logo" />
      </a>
    </div>
    <div class="appname">{{ t("appname") }}</div>
    <div class="by">By jinyaoMa</div>
  </my-container>
</template>

<style lang="scss" scoped>
.home {
  text-align: center;
  justify-content: center;
}
.top {
  text-align: center;
  font-size: 2em;
}
.banner {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}
.logo {
  height: 6em;
  padding: 0 3em 1em;
  will-change: filter;
  transition: filter 0.2s;

  &:hover {
    filter: drop-shadow(0 0 2em #46adeaaa);
  }
  &.wails:hover {
    filter: drop-shadow(0 0 1em #e55555aa);
  }
}
.divider {
  height: 3em;
  margin-bottom: 2.3em;
  border-left: var(--my-border-width) solid var(--my-color-border-light);
}
.appname {
  text-align: center;
  font-size: 3em;
}
.by {
  text-align: center;
  font-size: 1.3em;
  color: var(--my-color-text-regular);
}
</style>
