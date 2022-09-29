<script setup lang="ts">
import { ref } from "vue";
import { useI18n } from "vue-i18n";
import { EventsOn } from "../../wailsjs/runtime";
import { BrowserOpenURL } from "../../wailsjs/runtime";

const { t } = useI18n();

const isWebServiceStart = ref(false);
EventsOn("onWebServiceChanged", (isStart: boolean) => {
  console.log("onWebServiceChanged", isStart);
  isWebServiceStart.value = isStart;
});

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
      <template v-if="isWebServiceStart">
        <div class="sidebar-footer-line">
          <my-link @click="openVitePress" underline>
            <my-icon name="external-link"></my-icon>
            <span>{{ t("footer.openVitePress") }}</span>
          </my-link>
        </div>
        <div class="sidebar-footer-line">
          <my-link @click="openSwagger" underline>
            <my-icon name="external-link"></my-icon>
            <span>{{ t("footer.openSwagger") }}</span>
          </my-link>
        </div>
      </template>
      <div class="copyright">© 2022 jinyaoMa</div>
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
  &-footer-line {
    margin: 0.5em 0;
  }
}
.copyright {
  color: var(--my-color-text-placeholder);
  margin-top: 1em;
}
[class*="my-icon-"] {
  margin-right: 0.5em;
}
</style>
