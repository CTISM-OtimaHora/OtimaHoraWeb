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


function append_to_select_field(all_objs_arr, allowed_ids, select_div, option_tipo) {
    all_objs_arr.filter((obj) => allowed_ids.includes(obj.Id)).forEach((obj) => select_div.appendChild(new_option_child(obj, option_tipo)))
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
        const selected = profs.options[profs.selectedIndex]
        const dis_div = document.getElementById("disciplinas")
        
        if (selected.value == "default") {
            dis_div.innerHTML = ""
            dis_div.appendChild(new_option_child({Id: "default", Nome: "Selecione uma Disciplina"}, ""))
            append_to_select_field(s.Disciplinas, s.Disciplinas.map(d => d.Id), dis_div, "disciplina")
            return
        }

        if (dis_div.options[dis_div.selectedIndex].value !== "default") {
            return
        }

        const select_prof = s.Professores.filter((e) =>e.Id == selected.value)[0]

        const def = dis_div.firstElementChild
        dis_div.innerHTML = ""
        dis_div.appendChild(def)
        append_to_select_field(s.Disciplinas, select_prof.Disciplinas_ids.map(d => parseInt(d)), dis_div, "disciplina")
    })

    const dis_list = document.getElementById("disciplinas")
    dis_list.addEventListener("change", () => {

        const selected_dis_div = dis_list.options[dis_list.selectedIndex] 
        if (selected_dis_div.value == "default") {
            const f = profs.firstElementChild
            profs.innerHTML = ""
            profs.appendChild(f)
            append_to_select_field(s.Professores, s.Professores.map(p => p.Id), profs, "professor")
            return
        }

        if (profs.options[profs.selectedIndex].value !== "default") {
            return
        }

        const valid_profs = s.Professores.filter(p => p.Disciplinas_ids.includes(parseInt(selected_dis_div.value))).map(p => p.Id)

        const f = profs.firstElementChild
        profs.innerHTML = ""
        profs.appendChild(f)
        append_to_select_field(s.Professores, valid_profs, profs, "professor")
    })

    const cursos = document.getElementById("cursos")
    cursos.addEventListener("change", () => {
        const selected_curso_div = cursos.options[cursos.selectedIndex] 
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

            dis_list.innerHTML = ""
            dis_list.appendChild(new_option_child({Id: "default", Nome: "Selecione uma Disciplina"}, ""))
            append_to_select_field(s.Disciplinas, s.Disciplinas.map(d => d.Id), dis_list, "disciplina")

            return 
        }
        turmas.appendChild(def)

        const selected_curso = s.Cursos.filter((c) => c.Id == parseInt(selected_curso_div.value))[0]
        selected_curso.Etapas
            .forEach((etapa) => etapa
            .forEach((turma) => {turmas.appendChild(new_option_child(turma, "turma"))}))


        dis_list.innerHTML = ""
        dis_list.appendChild(new_option_child({Id: "default", Nome: "Selecione uma Disciplina"}, ""))
        append_to_select_field(s.Disciplinas, Object.keys(selected_curso.Curriculo).map(k => parseInt(k)), dis_list, "disciplina")
    })
});
