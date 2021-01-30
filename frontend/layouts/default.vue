<template>
  <div class=" min-h-screen h-full grid grid-rows-layout grid-cols-layout">
    <header class=" col-span-3 row-span-1 bg-green-700">
      <div class="flex justify-center items-center px-4 h-full">
        <div class="flex items-center w-1/5">
          <h2>Logo</h2>
        </div>
        <div class=" w-3/5">
          <!-- search bar-->
        </div>
        <div class=" w-1/5">
          <!-- left -->
        </div>
      </div>
    </header>
    <aside class="col-span-1 row-span-1 bg-gray-800 px-4 pt-2">
      <ul class="text-gray-300 m-0">
        <li class="py-2 text-center border-b border-gray-700">
          <nuxt-link to="estimates">
            Estimate</a>
          </nuxt-link>
        </li>
        <li class="py-2 text-center border-b border-gray-700">
          <nuxt-link to="settings">
            Settings</a>
          </nuxt-link>
        </li>
      </ul>
    </aside>
    <main class="col-span-1 pl-4 pt-4">
      <Nuxt />
    </main>
    <aside class="col-span-1 row-span-1 bg-gray-400 px-4 pt-2">
      <h3 class="text-xl text-gray-600 pb-2 mp-2 border-b border-gray-500 text-center">
        Properties
      </h3>
    </aside>
    <footer class=" col-span-3 flex justify-center items-center">
      Copyright Cunis Inc
    </footer>
  </div>
</template>

<style>
</style>
