function adicionar_disciplina() {
    const name = document.getElementById("new_disciplina").value;
    document.getElementById("new_disciplina").value = "";
    const disciplina = { id: undefined, nome: name };

    fetch(`http://localhost:3000/add-disciplina`, {
        method: "POST",
        credentials: "include",
        body: JSON.stringify(disciplina)
    })
    .then(response => response.text())
    .then(disciplina_id => {
        document.getElementById("disciplinas").appendChild(new_disciplina(name, disciplina_id));
    });
}

function new_disciplina(name, disciplina_id) {
    const disciplina = document.createElement("div");
    disciplina.classList.add("item");

    const p_name = document.createElement("div");
    p_name.textContent = name;
    disciplina.appendChild(p_name);

    const btn = document.createElement("button");
    btn.onclick = () => {
        window.location.replace(`/dashboard.html?tipo=disciplina&disciplina=${name}&id=${disciplina_id}`);
    };
    btn.textContent = "Editar";
    disciplina.appendChild(btn);

    return disciplina;
}

document.addEventListener("DOMContentLoaded", async () => {
    const res = await fetch(`http://localhost:3000/session`, {
        credentials: "include"
    });
    const s = await res.json();

    for (const disciplina of s.Disciplinas) {
        document.getElementById("disciplinas").appendChild(new_disciplina(
            disciplina.Nome,
            disciplina.Id
        ));
    }
});

