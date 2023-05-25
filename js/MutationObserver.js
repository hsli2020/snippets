<script>
  const queryParams = new URLSearchParams(window.location.search)
  const form = document.querySelector(".formkit-sticky-bar")
  if (queryParams.get("fromNewsletter") === "true") {
    window.location = window.location.pathname
    localStorage.setItem('subscribed-to-newsletter', true)
  }

  if (queryParams.get("fromNewsletter") === "false") {
    window.location = window.location.pathname
    localStorage.removeItem('subscribed-to-newsletter')
  }

  if (
    form == null &&
    JSON.parse(localStorage.getItem("subscribed-to-newsletter")) !== true &&
    JSON.parse(sessionStorage.getItem("form-closed")) !== true
  ) {
    const script = document.createElement("script")
    script.src = "https://web-dev-simplified.ck.page/23989b36d2/index.js"
    script.async = true
    script.dataset.uid = '23989b36d2'
   
    const observer = new MutationObserver((entries) => {
      entries.forEach(entry => {
        const formElem = [...entry.addedNodes].find(node => {
          if (node.matches == null) return
          return node.matches(".formkit-sticky-bar")
        })
        if (formElem == null) return

        formElem.addEventListener("transitionend", () => {
          if (formElem.dataset.active == null || formElem.dataset.active === "false") {
            formElem.remove()
          }
        })

        document.body.prepend(formElem)
        observer.disconnect()
      })
    })

    observer.observe(document.body, { childList: true })

    document.body.append(script)
  }

  document.addEventListener("click", e => {
    if (e.target.matches("[data-formkit-close]")) {
      sessionStorage.setItem("form-closed", true)
    }
  })
</script>
