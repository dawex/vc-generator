<template>
  <header class="fixed w-full z-30">
    <div class="absolute inset-0 bg-white bg-opacity-70 border-b border-slate-200 backdrop-blur -z-10 dark:bg-slate-900 dark:border-slate-800" aria-hidden="true"></div>
    <div class="mx-auto px-4 sm:px-6">
      <div class="flex items-center justify-between h-16 md:h-20">

        <!-- Site branding -->
        <div class="grow">
          <div class="flex items-center">
            <!-- Logo -->
            <router-link to="/" aria-label="API guide">
              <img src="../images/logo.png" width="160" height="32" alt="Docs" :class="darkMode === true ? 'image-inverted' : ''"/>
            </router-link>
            <!-- Search -->
            <div class="grow ml-4 md:ml-16 mt-3 hidden md:block">
              <h1 class="h4 text-slate-800 mb-4 dark:text-slate-200">Verification Plugin - API</h1>
            </div>
          </div>
        </div>

        <!-- Desktop navigation -->
        <nav class="flex">

          <!-- Right side elements links -->
          <ul class="flex grow justify-end flex-wrap items-center">
            <!-- <li class="ml-4">
              <a class="btn-sm inline-flex items-center text-slate-100 bg-blue-600 hover:bg-blue-700 shadow-sm" href="#0">Support</a>
            </li> -->
            <!-- Lights switch -->
            <!-- <li>
              <div class="flex flex-col justify-center ml-3">
                <input type="checkbox" name="light-switch" id="light-switch" v-model="darkMode" class="light-switch sr-only" />
                <label class="relative cursor-pointer p-2" for="light-switch">
                  <svg class="dark:hidden" width="16" height="16" xmlns="http://www.w3.org/2000/svg">
                    <path class="fill-blue-400" d="M7 0h2v2H7zM12.88 1.637l1.414 1.415-1.415 1.413-1.413-1.414zM14 7h2v2h-2zM12.95 14.433l-1.414-1.413 1.413-1.415 1.415 1.414zM7 14h2v2H7zM2.98 14.364l-1.413-1.415 1.414-1.414 1.414 1.415zM0 7h2v2H0zM3.05 1.706 4.463 3.12 3.05 4.535 1.636 3.12z" />
                    <path class="fill-blue-500" d="M8 4C5.8 4 4 5.8 4 8s1.8 4 4 4 4-1.8 4-4-1.8-4-4-4Z" />
                  </svg>
                  <svg class="hidden dark:block" width="16" height="16" xmlns="http://www.w3.org/2000/svg">
                    <path class="fill-blue-500" d="M6.2 1C3.2 1.8 1 4.6 1 7.9 1 11.8 4.2 15 8.1 15c3.3 0 6-2.2 6.9-5.2C9.7 11.2 4.8 6.3 6.2 1Z" />
                    <path class="fill-blue-500" d="M12.5 5a.625.625 0 0 1-.625-.625 1.252 1.252 0 0 0-1.25-1.25.625.625 0 1 1 0-1.25 1.252 1.252 0 0 0 1.25-1.25.625.625 0 1 1 1.25 0c.001.69.56 1.249 1.25 1.25a.625.625 0 1 1 0 1.25c-.69.001-1.249.56-1.25 1.25A.625.625 0 0 1 12.5 5Z" />
                  </svg>
                  <span class="sr-only">Switch to light / dark version</span>
                </label>
              </div>
            </li> -->
          </ul>

        </nav>

      </div>
    </div>
  </header>
</template>

<script>
import { ref, watch } from 'vue'

export default {
  name: 'Header',
  setup() {

    const darkMode = ref(localStorage.getItem('dark-mode'))

    watch(darkMode, () => {
      localStorage.setItem('dark-mode', darkMode.value)
      if (darkMode.value) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
    })

    return {
      darkMode,
    }
  },
}
</script>

<style>
.image-inverted {
  filter: invert(100%);
}
</style>