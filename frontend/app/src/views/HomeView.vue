<script setup lang="ts">
import { RouterLink } from 'vue-router'
import HelloWorld from '../components/HelloWorld.vue'
import TheWelcome from '../components/TheWelcome.vue'
import { Events } from '@wailsio/runtime'
import { api } from '@jinyaoma/my-sdk'
import { GreetService } from '@jinyaoma/my-sdk/bindings/majinyao.cn/my-app/backend/cmd/wails/services'

api.auth
  .authLogin({
    account: '',
    password: ''
  })
  .then((res) => {
    console.log(res)
  })

const doGreet = async () => {
  const name = (document.getElementById('name')! as HTMLInputElement).value
  const result = document.getElementById('result')
  result!.innerText = await GreetService.Greet(name)
}

Events.On('time', (time: { data: string }) => {
  const timeElement = document.getElementById('time')
  timeElement!.innerText = time.data
})
</script>

<template>
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav>
        <RouterLink to="/">Home</RouterLink>
      </nav>
    </div>
  </header>

  <main>
    <TheWelcome />
    <div class="container">
      <div>
        <a wml-openURL="https://wails.io" href="https://wails.io" target="_blank">
          <img src="/wails.png" class="logo" alt="Wails logo" />
        </a>
        <a
          wml-openURL="https://developer.mozilla.org/en-US/docs/Web/JavaScript"
          href="https://developer.mozilla.org/en-US/docs/Web/JavaScript"
          target="_blank"
        >
          <img src="/javascript.svg" class="logo vanilla" alt="JavaScript logo" />
        </a>
      </div>
      <h1>Wails + Javascript</h1>
      <div class="card">
        <div class="result" id="result">Please enter your name below 👇</div>
        <div class="input-box" id="input">
          <input class="input" id="name" type="text" autocomplete="off" />
          <button class="btn" @click="doGreet">Greet</button>
        </div>
      </div>
      <div class="footer">
        <div>
          <p>Click on the Wails logo to learn more</p>
        </div>
        <div>
          <p id="time">Listening for Time event...</p>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
