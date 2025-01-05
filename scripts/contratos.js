function adicionar_contrato() {
    const profs = document.getElementById('professores');
    const dis = document.getElementById('disciplinas');
    // CURSO NÂO TEM DISPO // const cursos = document.getElementById('cursos');
    const turmas = document.getElementById('turmas');
    const recs = document.getElementById('recursos');

    let selected = [
        profs.options[profs.selectedIndex],
        dis.options[dis.selectedIndex],
        turmas.options[turmas.selectedIndex],
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

function new_contrato(id, participantes, tipos) {
    const contrato = document.createElement("div")
    let string  = ""
    for (const [idx, i] of participantes.entries()) {
        string += ` ${tipos[idx]}: ${i.Nome} + `
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


function append_to_select_field(all_objs_arr, session, select_div, option_tipo) {
    // const cursos_div = document.getElementById("cursos")
    const turmas_div = document.getElementById("turmas")
    const profs_div = document.getElementById("professores")
    const disci_div = document.getElementById("disciplinas")

    let allowed_ids = all_objs_arr.map(o => o.Id);

    if (option_tipo === "professor") {
        const sel_dis_div = disci_div.options[disci_div.selectedIndex]
        const sel_tur_div = turmas_div.options[turmas_div.selectedIndex]
        
        if (sel_tur_div.value !== "default") {
            let etapa = undefined;
            for (const curso  of session.Cursos) {
                for (const et of curso.Etapas) {
                    if (et.Turmas.map(t => t.Id).includes(parseInt(sel_tur_div.value))) {
                        etapa = et
                    }
                }
            }

            console.log(etapa)
            allowed_ids = []
            for (const dis_id of Object.keys(etapa.Curriculo).map(k => parseInt(k))) {
                allowed_ids = allowed_ids.concat(
                    all_objs_arr
                    .filter(p => p.Disciplinas_ids.includes(dis_id))
                    .map(p => p.Id))
            }
        } 
        if (sel_dis_div.value !== "default") {
            allowed_ids = all_objs_arr.filter(p => p.Disciplinas_ids.includes(parseInt(sel_dis_div.value))).map(p => p.Id)
        } 
    }

    if (option_tipo === "disciplina") {
        const sel_tur_div = turmas_div.options[turmas_div.selectedIndex]
        const sel_pro_div = profs_div.options[profs_div.selectedIndex]

        if (sel_pro_div.value !== "default") {
            const pro = session.Professores.filter(p => p.Id === parseInt(sel_pro_div.value))[0]
            allowed_ids = all_objs_arr.filter(d => pro.Disciplinas_ids.includes(d.Id)).map(p => p.Id)
        } 

        if (sel_tur_div.value !== "default") {
            let etapa = undefined;
            for (const curso  of session.Cursos) {
                for (const et of curso.Etapas) {
                    if (et.Turmas.map(t => t.Id).includes(parseInt(sel_tur_div.value))) {
                        etapa = et
                    }
                }
            }
            console.log(etapa)
            allowed_ids = Object.keys(etapa.Curriculo).map(k => parseInt(k))
            console.log(allowed_ids)

        }
    }

    if (option_tipo === "curso") {
        // const sel_dis_div = disci_div.options[disci_div.selectedIndex]
        // const sel_pro_div = profs_div.options[profs_div.selectedIndex]
        //
        // if (sel_pro_div.value !== "default") {
        //     const pro = session.Professores.filter(p => p.Id === parseInt(sel_pro_div.value))[0]
        //     allowed_ids = []
        //     for (dis_id of pro.Disciplinas_ids) {
        //         dis_id = dis_id.toString()
        //         allowed_ids = allowed_ids.concat(all_objs_arr.filter(c => Object.keys(c.Curriculo).includes(dis_id)).map(c => c.Id))
        //     }
        // }
        // if (sel_dis_div.value !== "default") {
        //     console.log("lol")
        //     allowed_ids = all_objs_arr.filter(c => Object.keys(c.Curriculo).includes(sel_dis_div.value)).map(c=>c.Id)
        // }
    }

    all_objs_arr
        .filter((obj) => allowed_ids.includes(obj.Id))
        .forEach((obj) => select_div.appendChild(new_option_child(obj, option_tipo)))
} 

document.addEventListener("DOMContentLoaded", async () => {
    const cursos_div = document.getElementById("cursos")
    const profs_div = document.getElementById("professores")
    const turmas_div = document.getElementById("turmas")
    const disci_div = document.getElementById("disciplinas")
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
            document.getElementById("contratos").appendChild(new_contrato(c.Id, c.Participantes, c.Tipo_por_participante))
        }
    }

    const profs = document.getElementById("professores")
    profs.addEventListener("change", () => {
        if (disci_div.options[disci_div.selectedIndex].value === "default") {
            const def = disci_div.firstElementChild
            disci_div.innerHTML = ""
            disci_div.appendChild(def)
            append_to_select_field(s.Disciplinas, s, disci_div, "disciplina")
        }
    })

    
    disci_div.addEventListener("change", () => {
        const f = profs.firstElementChild
        profs.innerHTML = ""
        profs.appendChild(f)
        append_to_select_field(s.Professores, s, profs, "professor")
    })

    turmas_div.addEventListener("change", () => {
        const f1 = disci_div.firstElementChild
        disci_div.innerHTML = ""
        disci_div.appendChild(f1)
        append_to_select_field(s.Disciplinas, s, disci_div, "disciplina")

        const f = profs.firstElementChild
        profs.innerHTML = ""
        profs.appendChild(f)
        append_to_select_field(s.Professores, s, profs, "professor")
    })

    cursos_div.addEventListener("change", () => {
        const selected_curso_div = cursos_div.options[cursos_div.selectedIndex] 
        const def = document.createElement("option")     
        def.textContent = "Selecione uma turma"
        def.value = "default"

        const turmas = document.getElementById("turmas")
        turmas.innerHTML = ""
        if (selected_curso_div.value === "default") {
            const def = document.createElement("option")     
            def.textContent = "Selecione uma turma (Selecione um Curso primeiro)"
            def.value = "default"
            turmas.appendChild(def)
        } else {
            turmas.appendChild(def)
            const selected_curso = s.Cursos.filter((c) => c.Id == parseInt(selected_curso_div.value))[0]
            selected_curso.Etapas
                .forEach((etapa) => etapa.Turmas
                    .forEach((turma) => {turmas.appendChild(new_option_child(turma, "turma"))}))
        }

        const f1 = disci_div.firstElementChild
        disci_div.innerHTML = ""
        disci_div.appendChild(f1)
        append_to_select_field(s.Disciplinas, s, disci_div, "disciplina")

        const f = profs.firstElementChild
        profs.innerHTML = ""
        profs.appendChild(f)
        append_to_select_field(s.Professores, s, profs, "professor")

    })
});
