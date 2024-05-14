function adicionar_contrato() {
    const profs = document.getElementById('professores');
    const dis = document.getElementById('disciplinas');
    const cursos = document.getElementById('cursos');
    const recs = document.getElementById('recursos');

    let selected = [
        profs.options[profs.selectedIndex],
        dis.options[dis.selectedIndex],
        cursos.options[cursos.selectedIndex],
        recs.options[recs.selectedIndex],
    ]

    selected = selected
        .filter((e) => e && e.value != "default")
        .map(
            (e) => {return {
                id: parseInt(e.value), nome: e.textContent, tipo: e.alt
            }});


    fetch("http://localhost:3000/session/add/contrato", {
        credentials: "include",
        method: "POST",
        body: JSON.stringify(selected)
    }).then(alert("saved")).then(window.location.reload())
}

function new_contrato(id, participantes) {
    const contrato = document.createElement("div")
    let string  = ""
    for (const i of participantes) {
        string += ` ${i.Tipo}: ${i.Nome} + `
    }
    contrato.textContent = string.slice(0, -3)

    const bttn = document.createElement("a")
    bttn.classList.add("add")
    bttn.textContent = "Ver"
    bttn.onclick = () => {
        window.location.replace(`/dashboard.html?tipo=contrato&id=${id}`);
    }
    contrato.appendChild(bttn)

    return contrato
}

document.addEventListener("DOMContentLoaded", async () => {
    const res = await fetch(`http://localhost:3000/session`, {
        credentials: "include"
    });
    const s = await res.json()
    
    if (s.Professores) {
        for (const professor of s.Professores) {
            const child = document.createElement("option")
            child.value = professor.Id
            child.textContent = professor.Nome
            child.alt = "professor"

            document.getElementById("professores").appendChild(child);
        }
    }

    if (s.Disciplinas) {
        for (const disciplina of s.Disciplinas) {
            const child = document.createElement("option")
            child.value = disciplina.Id
            child.textContent = disciplina.Nome
            child.alt = "disciplina"

            document.getElementById("disciplinas").appendChild(child);
        }
    }

    if (s.Cursos) {
        for (const curso of s.Cursos) {
            const child = document.createElement("option")
            child.value = curso.Id
            child.textContent = curso.Nome
            child.alt = "curso"

            document.getElementById("cursos").appendChild(child);
        }
    }
    if (s.Recursos) {
        for (const recurso of s.Recursos) {
            const child = document.createElement("option")
            child.value = recurso.Id
            child.textContent = recurso.Nome
            child.alt = "recurso"

            document.getElementById("recursos").appendChild(child);
        }
    }


    if (s.Contratos) {
        for (const c of s.Contratos) {
            console.log(c)
            document.getElementById("contratos").appendChild(new_contrato(c.Id, c.Participantes))
        }
    }
});
