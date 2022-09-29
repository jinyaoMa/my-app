<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { useColorTheme } from "../store/color-theme";
import { EventsOn } from "../../wailsjs/runtime";
import { GetOptions } from "../../wailsjs/go/service/settings";
import { ref } from "vue";

const { t, locale } = useI18n();
const colorTheme = useColorTheme();

const options = ref<app.Config>();
GetOptions().then((config) => {
  options.value = config;
});
</script>

<template>
  <my-container class="settings">
    <my-main>
      <h2>{{ t("menu.main.settings") }}</h2>
      <my-form>
        <my-form-group :legend="t('settings.general.legend')">
          <my-form-item
            :label="t('settings.general.displayLanguage')"
            label-width="10em"
          >
            {{ t("lang") }}
          </my-form-item>
          <my-form-item
            :label="t('settings.general.colorTheme')"
            label-width="10em"
          >
            {{ t(`theme.${colorTheme.theme}`) }}
          </my-form-item>
          <my-form-item
            :label="t('settings.general.logPath')"
            label-width="10em"
          >
            {{ options?.LogPath }}
          </my-form-item>
        </my-form-group>
        <my-form-group :legend="t('settings.webService.legend')">
          <my-form-item
            :label="t('settings.webService.portHttp')"
            label-width="10em"
          >
            {{ options?.Web.PortHttp }}
          </my-form-item>
          <my-form-item
            :label="t('settings.webService.portHttps')"
            label-width="10em"
          >
            {{ options?.Web.PortHttps }}
          </my-form-item>
          <my-form-item
            :label="t('settings.webService.dirCerts')"
            label-width="10em"
          >
            {{ options?.Web.DirCerts }}
          </my-form-item>
        </my-form-group>
      </my-form>
    </my-main>
  </my-container>
</template>

<style lang="scss" scoped>
.settings {
  width: 100%;
  max-width: 560px;
}
h2 {
  font-size: 2em;
}
</style>
