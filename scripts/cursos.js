function adicionar_curso() {
    const name = document.getElementById("new_curso").value;
    document.getElementById("new_curso").value = "";
    curso = {id: undefined, nome: name, turmas: []}
    fetch("http://localhost:3000/add-curso", {method:"POST", credentials:"include", body:JSON.stringify(curso)})
        .then(response => response.text()) // Get the plain text ID from the response
        .then(id => {
            document.getElementById("cursos").appendChild(new_curso(name, id));
        });
}

function new_curso(name, id) {
    let curso = document.createElement("div");

    let c_name = document.createElement("div");
    c_name.textContent = name;
    curso.appendChild(c_name);

    let btn = document.createElement("button");
    btn.onclick = () => {window.location.replace(`/dashboard.html?tipo=curso&id=${id}&curso=${name}`)}
    btn.textContent = "turmas";
    curso.appendChild(btn);

    curso.classList.add("item");
    curso.classList.add("curso");
    return curso
}

document.addEventListener('DOMContentLoaded', async function() {
    const res = await fetch("http://localhost:3000/session", {credentials:"include"})
    const obj = await res.json()

    console.log(obj.Cursos);
    for (c of obj.Cursos) {
        document.getElementById("cursos").appendChild(new_curso(c.Nome, c.Id));
    }
});

