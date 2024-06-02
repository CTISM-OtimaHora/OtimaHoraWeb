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


    fetch("http://localhost:3000/contrato/add", {
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
        window.location.replace(`/OtimaHoraWeb/dashboard.html?tipo=contrato&id=${id}`);
    }
    contrato.appendChild(bttn)

    return contrato
}

function new_option_child(child_obj, tipo) {
    const child = document.createElement("option")
    child.value = child_obj.Id
    child.textContent = child_obj.Nome
    child.alt = tipo
    return child
}

document.addEventListener("DOMContentLoaded", async () => {
    const res = await fetch(`http://localhost:3000/session`, {
        credentials: "include"
    });
    const s = await res.json()

    
    if (s.Professores) {
        for (const professor of s.Professores) {
            document.getElementById("professores").appendChild(new_option_child(professor, "professor"));
        }
    }

    if (s.Disciplinas) {
        for (const disciplina of s.Disciplinas) {
            document.getElementById("disciplinas").appendChild(new_option_child(disciplina, "disciplina"));
        }
    }

    if (s.Cursos) {
        for (const curso of s.Cursos) {
            document.getElementById("cursos").appendChild(new_option_child(curso, "curso"));
        }
    }
    if (s.Recursos) {
        for (const recurso of s.Recursos) {
            document.getElementById("recursos").appendChild(new_option_child(recurso, "recurso"))
        }
    }


    if (s.Contratos) {
        for (const c of s.Contratos) {
            document.getElementById("contratos").appendChild(new_contrato(c.Id, c.Participantes))
        }
    }

    const profs = document.getElementById("professores")
    profs.addEventListener("change", () => {
        if (!s.Professores || !s.Disciplinas) {
            return;
        }
        const dis_div = document.getElementById("disciplinas")

        const selected = profs.options[profs.selectedIndex]
        let valid_disciplinas = s.Disciplinas

        if (selected.value !== "default") {
            const select_prof = s.Professores.filter((e) =>e.Id == selected.value)[0]

            if (select_prof.Disciplinas_ids) {
                valid_disciplinas = s.Disciplinas.filter((d) => select_prof.Disciplinas_ids.includes(d.Id))
            }
            
            const selected_dis = dis_div.options[dis_div.selectedIndex]
            if (selected_dis.value !== "default") {
                if (select_prof.Disciplinas_ids.includes(parseInt(selected_dis.value))) {
                    return
                }
            }
        }

        // manter a opção "selecionar"
        const def = dis_div.firstElementChild
        dis_div.innerHTML = ""
        dis_div.appendChild(def)

        for (const disciplina of valid_disciplinas) {
            document.getElementById("disciplinas").appendChild(new_option_child(disciplina, "disciplina"));
        }
    })

    const dis_list = document.getElementById("disciplinas")
    dis_list.addEventListener("change", () => {
        const selected_dis_div = dis_list.options[dis_list.selectedIndex] 
        let valid_profs = s.Professores

        if (selected_dis_div.value != "default") {
            valid_profs = s.Professores.filter((p) => p.Disciplinas_ids !== null).filter((p) => p.Disciplinas_ids.includes(parseInt(selected_dis_div.value)))

            const selected_prof = profs.options[profs.selectedIndex]
            if (selected_prof.value !== "default") {
                if (valid_profs.map((p) => p.Id).includes(parseInt(selected_prof.value))) {
                    return
                }
            }
        }

        
        const def = profs.firstElementChild
        profs.innerHTML = ""
        profs.appendChild(def)

        if (valid_profs.lenght == 0) {
            profs.innerHTML = ""
            const bad = document.createElement("option")
            bad.textContent = "Nenhuma professor pode lecionar essa disciplina"
            profs.appendChild(bad)
            return
        }

        for (const professor of valid_profs) {
            profs.appendChild(new_option_child(professor, "professor"))
        }

    })
});
