document.addEventListener('DOMContentLoaded', async function() {
    await fetch("http://localhost:3000/add-session", {credentials:"include"})

    const res = await fetch("http://localhost:3000/session", {credentials:"include"})
});
