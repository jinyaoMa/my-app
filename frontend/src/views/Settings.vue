<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { useColorTheme } from "../store/color-theme";
import { EventsOn } from "../../wailsjs/runtime";
import {
  GetOptions,
  ChooseLogPath,
  ChooseDirCerts,
} from "../../wailsjs/go/service/settings";
import { ChangeLanguage, ChangeColorTheme } from "../../wailsjs/go/tray/tray";
import { ResponseState } from "../../packages/components/types";
import { ref } from "vue";

const { t, locale, availableLocales } = useI18n();
const colorTheme = useColorTheme();

const options = ref<app.Config>();
GetOptions().then((config) => {
  console.log(config);
  options.value = config;
});

const changeDisplayLanguage = (newLang: string) => {
  ChangeLanguage(newLang);
};
const changeColorTheme = (newTheme: string) => {
  ChangeColorTheme(newTheme);
};

const changeLogPath = async (res: (state: ResponseState) => void) => {
  if (options.value) {
    const newLogPath = await ChooseLogPath(
      options.value.LogPath,
      t("settings.chooseLogPath")
    );
    if (newLogPath === "") {
      res("warning");
      return;
    }
    options.value!.LogPath = newLogPath;
    res("success");
  }
};
const changeDirCerts = async (res: (state: ResponseState) => void) => {
  if (options.value) {
    const newDirCerts = await ChooseDirCerts(
      options.value!.Web.DirCerts,
      t("settings.chooseDirCerts")
    );
    if (newDirCerts === "") {
      res("warning");
      return;
    }
    options.value!.Web.DirCerts = newDirCerts;
    res("success");
  }
};
</script>

<template>
  <my-container class="settings">
    <my-main>
      <h2>{{ t("menu.main.settings") }}</h2>
      <my-form>
        <my-form-group
          :legend="t('settings.general.legend')"
          :label-width="locale === 'zh' ? '7em' : '9em'"
        >
          <my-form-item :label="t('settings.general.displayLanguage')">
            <my-select
              name="DisplayLanguage"
              :display-value="locale"
              @change="changeDisplayLanguage"
            >
              <my-option
                v-for="lang in availableLocales"
                :key="lang"
                :value="lang"
              >
                {{ t(`lang.${lang}`) }}
              </my-option>
            </my-select>
          </my-form-item>
          <my-form-item :label="t('settings.general.colorTheme')">
            <my-select
              name="ColorTheme"
              :display-value="colorTheme.theme"
              @change="changeColorTheme"
            >
              <my-option value="system">
                {{ t(`theme.system`) }}
              </my-option>
              <my-option value="light">
                {{ t(`theme.light`) }}
              </my-option>
              <my-option value="dark">
                {{ t(`theme.dark`) }}
              </my-option>
            </my-select>
          </my-form-item>
          <my-form-item :label="t('settings.general.logPath')">
            <my-input name="LogPath" :display-value="options?.LogPath" disabled>
              <template #append>
                <my-button type="primary" @click="changeLogPath">
                  {{ t("settings.choosePath") }}
                </my-button>
              </template>
            </my-input>
          </my-form-item>
        </my-form-group>
        <my-form-group
          :legend="t('settings.webService.legend')"
          :label-width="locale === 'zh' ? '7em' : '10em'"
        >
          <my-form-item :label="t('settings.webService.autoStart')">
            {{ options?.Web.AutoStart }}
          </my-form-item>
          <my-form-item :label="t('settings.webService.portHttp')">
            {{ options?.Web.PortHttp }}
          </my-form-item>
          <my-form-item :label="t('settings.webService.portHttps')">
            {{ options?.Web.PortHttps }}
          </my-form-item>
          <my-form-item :label="t('settings.webService.dirCerts')">
            <my-input
              name="DirCerts"
              :display-value="options?.Web.DirCerts"
              disabled
            >
              <template #append>
                <my-button type="primary" @click="changeDirCerts">
                  {{ t("settings.choosePath") }}
                </my-button>
              </template>
            </my-input>
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
  margin-top: 0;
}
</style>
