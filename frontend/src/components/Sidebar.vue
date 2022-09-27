<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { BrowserOpenURL } from "../../wailsjs/runtime";

const { t } = useI18n();

const openVitePress = () => {
  BrowserOpenURL("https://localhost:10443/docs/");
};
const openSwagger = () => {
  BrowserOpenURL("https://localhost:10443/swagger/index.html");
};
</script>
<template>
  <my-container class="sidebar" height="100%">
    <my-main class="sidebar-main">
      <my-menu type="primary">
        <my-menu-group :title="t('menu.main.title')">
          <my-menu-item to="/">
            <my-icon name="home" fix-width></my-icon>
            <span>{{ t("menu.main.home") }}</span>
          </my-menu-item>
          <my-menu-item to="/settings">
            <my-icon name="settings" fix-width></my-icon>
            <span>{{ t("menu.main.settings") }}</span>
          </my-menu-item>
          <my-menu-item to="/about">
            <my-icon name="info" fix-width></my-icon>
            <span>{{ t("menu.main.about") }}</span>
          </my-menu-item>
        </my-menu-group>
        <my-menu-group :title="t('menu.test.title')">
          <my-menu-item to="/test">
            <my-icon name="home" fix-width></my-icon>
            <span>{{ t("menu.test.test") }}</span>
          </my-menu-item>
        </my-menu-group>
      </my-menu>
    </my-main>
    <my-footer class="sidebar-footer">
      <div>
        <my-link @click="openVitePress" underline>
          <my-icon name="external-link"></my-icon>
          <span>{{ t("footer.openVitePress") }}</span>
        </my-link>
      </div>
      <div>
        <my-link @click="openSwagger" underline>
          <my-icon name="external-link"></my-icon>
          <span>{{ t("footer.openSwagger") }}</span>
        </my-link>
      </div>
      <div
        :style="{
          color: 'var(--my-color-text-placeholder)',
          marginTop: '1em',
        }"
      >
        © 2022 jinyaoMa
      </div>
    </my-footer>
  </my-container>
</template>

<style lang="scss" scoped>
.sidebar {
  border-right: var(--my-border-width) solid var(--my-color-border-extra-light);
  &-main,
  &-footer {
    padding: 2em;
  }
}
[class*="my-icon-"] {
  margin-right: 0.5em;
}
</style>
