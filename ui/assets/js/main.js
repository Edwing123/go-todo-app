const formatter = new Intl.DateTimeFormat("default", {
    year: "numeric",
    month: "short",
    day: "2-digit",
    hour12: true,
    hour: "numeric",
    minute: "numeric",
    second: "numeric",
})

window.addEventListener("alpine:init", () => {
    Alpine.data("datetime", () => ({
        prettyDateTime() {
            return formatter.format(new Date(this.$root.getAttribute('datetime')))
        }
    }))
})
