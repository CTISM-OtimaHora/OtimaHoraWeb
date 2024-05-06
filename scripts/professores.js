function adicionar_professor() {
    const name = document.getElementById("new_professor").value;
    document.getElementById("new_professor").value = "";
    const professor = { id: undefined, nome: name };

    fetch(`http://localhost:3000/add-professor/`, {
        method: "POST",
        credentials: "include",
        body: JSON.stringify(professor)
    })
    .then(response => response.text())
    .then(professor_id => {
        document.getElementById("professores").appendChild(new_professor(name, professor_id));
    });
}

function new_professor(name, professor_id) {
    const professor = document.createElement("div");
    professor.classList.add("item");

    const p_name = document.createElement("div");
    p_name.textContent = name;
    professor.appendChild(p_name);

    const btn = document.createElement("button");
    btn.onclick = () => {
        window.location.replace(`/professor_dashboard.html?professor=${name}&professor_id=${professor_id}`);
    };
    btn.textContent = "Editar";
    professor.appendChild(btn);

    return professor;
}

document.addEventListener("DOMContentLoaded", async () => {
    const res = await fetch(`http://localhost:3000/session`, {
        credentials: "include"
    });
    const s = await res.json();

    for (const professor of s.Professores) {
        document.getElementById("professores").appendChild(new_professor(
            professor.Nome,
            professor.Id
        ));
    }
});

