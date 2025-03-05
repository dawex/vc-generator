<template>
  <div>
    <router-view />
  </div>
</template>

<script>
const scrollSpy = () => {
  const targets = document.querySelectorAll('[data-scrollspy-target]')
  const links = document.querySelectorAll('[data-scrollspy-link]')
  if (links.length < 1) return
  const addActive = (i) => {
    const link = links[i] ? links[i] : links[0]
    link.classList.add('scrollspy-active')
  }
  const removeActive = (i) => {
    links[i].classList.remove('scrollspy-active')
  }
  const removeAllActive = () => [...Array(targets.length).keys()].forEach((link) => removeActive(link))
  const targetMargin = 100
  let currentActive = 0
  addActive(0)
  // listen for scroll events
  window.addEventListener('scroll', () => {
    const current = targets.length - [...targets].reverse().findIndex((target) => window.scrollY >= target.offsetTop - targetMargin) - 1
    if (current !== currentActive) {
      removeAllActive()
      currentActive = current
      addActive(current)
    }
  })  
}

export default {
  mounted() {
    window.addEventListener('load', () => {
      scrollSpy()
    })
    if (this.$router) {
      this.$watch('$route', () => {
        this.$nextTick(() => {      
          scrollSpy()
        })
      })          
    }
  }
}
</script>
