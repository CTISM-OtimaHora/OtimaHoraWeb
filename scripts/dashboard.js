function toggleStatus(cell) {
   if (cell.classList.contains("sim")) {
      cell.classList.remove("sim");
      cell.classList.add("talvez");
      cell.textContent = "TALVEZ";
   } else if (cell.classList.contains("talvez")) {
      cell.classList.remove("talvez");
      cell.classList.add("nao");
      cell.textContent = "NAO";
   } else if (cell.classList.contains("nao")) {
      cell.classList.remove("nao");
      cell.classList.add("sim");
      cell.textContent = "SIM";
   } else {
      cell.classList.add("sim");
      cell.textContent = "SIM";
   }
}

const PERIODOS = 5
const DIAS = 5

function toggleStatusCol(col_idx) {
    for (let i = 0; i < PERIODOS; i ++) {
        toggleStatus(document.getElementById(`${i}-${col_idx}`));
    }
}
function toggleStatusRow(row_idx) {
    for (let i = 0; i < DIAS; i ++) {
        toggleStatus(document.getElementById(`${row_idx}-${i}`));
    }
}

function toggleAll() {
    for (let i = 0; i < PERIODOS; i++) {
        toggleStatusRow(i)
    }
}

function get_disp() {
    const mat = []

    for (let p_idx = 0; p_idx < PERIODOS; p_idx ++){
        const row = []
        for (let day_idx = 0; day_idx < DIAS; day_idx++) {
            const text = document.getElementById(`${p_idx}-${day_idx}`).textContent
            let value = 1

            if (text === "TALVEZ") {
                value = 0
            }
            if (text === "NAO") {
                value = -1
            }

            row.push(value)
        }
        mat.push(row)
    }

    return mat 
}

function set_disp(matrix) {
    for (let p_idx = 0; p_idx < PERIODOS; p_idx++) {
        for (let day_idx = 0; day_idx < DIAS; day_idx++) {
            const value = matrix[p_idx][day_idx]
            let text = "SIM"

            if (value == 0) {
                text = "TALVEZ"
            }
            if (value == -1) {
                text = "NAO"
            }

            const cell = document.getElementById(`${p_idx}-${day_idx}`)
            cell.textContent = text
            cell.classList.remove("sim")
            cell.classList.remove("talvez")
            cell.classList.remove("nao")
            cell.classList.add(text.toLowerCase())
        }
    }
}


function save_item(obj = undefined) {
    console.log(obj)
    const params = new URLSearchParams(window.location.search)
    let dispo = undefined

    // etapa e curso apagam a div com a dispo
    if (document.getElementById("dispo")) {
        dispo = get_disp()
    }

    if (obj === undefined) {
        obj = {
            Id: parseInt(params.get("id")),
            Nome: params.get(params.get("tipo")),
            Dispo: dispo,
        }
    }
    if (params.get("tipo") === "professor") {
        const dis_ids = [...document.querySelectorAll(".disciplina-check:checked")].map(e => parseInt(e.value))
        obj.Disciplinas_ids = dis_ids
    }
    obj.Dispo = dispo
    let url = `http://localhost:3000/${params.get("tipo")}/set/${params.get("id")}`
    if (params.get("tipo") === "turma") {
        fetch(`http://localhost:3000/curso/get/${params.get("curso_pai")}`, {credentials: "include"})
            .then(res => res.json().then(curso => {
                curso.Etapas[obj.Etapa_idx].Turmas[obj.Idx_in_etapa] = obj
                url = `http://localhost:3000/curso/set/${params.get("curso_pai")}`
                fetch(url, 
                    {
                        credentials: "include",
                        method: "PUT",
                        body: JSON.stringify(curso)
                    }).then(alert("saved"))
                
            }
            )
            )
        return
    }
    if (params.get("tipo") === "etapa") {
        fetch(`http://localhost:3000/curso/get/${params.get("curso_pai")}`, {credentials: "include"})
            .then(res => res.json().then((curso) => {
                curso.Etapas[obj.Idx_in_Curso] = obj
                url = `http://localhost:3000/curso/set/${params.get("curso_pai")}`
                fetch(url, 
                    {
                        credentials: "include",
                        method: "PUT",
                        body: JSON.stringify(curso)
                    }).then(alert("saved"))
                
            }
            )
            )
        return
    }
    fetch(url, 
        {
            credentials: "include",
            method: "PUT",
            body: JSON.stringify(obj)
        }).then(alert("saved"))

}

document.addEventListener('DOMContentLoaded', async function() {
    const params = new URLSearchParams(window.location.search)
    let disp;
    let obj;
    if (params.get("tipo") === "turma") {
        const res = await fetch(`http://localhost:3000/curso/get/${params.get("curso_pai")}`, {credentials:"include"})
        const curso = await res.json()
        obj = curso.Etapas[params.get("etapa_pai")].Turmas.filter(t => t.Id == params.get("id"))[0]
        disp = obj.Dispo
    } else if (params.get("tipo") === "etapa") {
        const res = await fetch(`http://localhost:3000/curso/get/${params.get("curso_pai")}`, {credentials:"include"})
        const curso = await res.json()
        console.log(curso)
        obj = curso.Etapas[params.get("id")]
    } else {
        const res = await fetch(`http://localhost:3000/${params.get("tipo")}/get/${params.get("id")}`, {credentials: "include"})
        obj = await res.json()
        disp = obj.Dispo
    }
    console.log(obj)


    if (params.get("tipo") == "contrato") {
        handle_contrato(obj)
    }

    if (params.get("tipo") == "professor") {
        await handle_professor(obj)
    }

    // parte de curso
    if (params.get("tipo") === "curso") {
        handle_curso(obj)
    }

    if (params.get("tipo") === "etapa") {
        await handle_etapa(obj)
    }

    // end etapa
    
    const nome_div = document.getElementById("nome")
    nome_div.textContent = obj.Nome

    document.getElementById("save").onclick = () => save_item(obj)
    nome_div.onclick = () => {

        const str = prompt("Insira o novo nome")
        if (!str || str === "" || str.length == 0 ) {
            return
        }
        nome.textContent = str
        obj.Nome = nome_div.textContent
        save_item(obj)
    }

    if (params.get("tipo") !== "etapa" && params.get("tipo") !== "curso") {
        set_disp(disp)
    }
})

function handle_contrato(obj) {
    const part = document.getElementById("adicional")

    for (const [i, p] of obj.Participantes.entries()) {
        console.log(p)
        const child = document.createElement("div")   

        child.textContent = `${obj.Tipo_por_participante[i]} - ${p.Nome}  `

        let url = ""
        if (obj.Tipo_por_participante[i] == "turma") {
            url = `/OtimaHoraWeb/dashboard.html?etapa_pai=${p.Etapa_idx}&curso_pai=${p.Curso_id}&tipo=${obj.Tipo_por_participante[i]}&${obj.Tipo_por_participante[i]}=${p.Nome}&id=${p.Id}`
        } else {
            url = `/OtimaHoraWeb/dashboard.html?tipo=${obj.Tipo_por_participante[i]}&${obj.Tipo_por_participante[i]}=${p.Nome}&id=${p.Id}`
        }
        child.onclick = () => {
            window.location.replace(url)
        }
        part.appendChild(child)
    }
    document.getElementById("save").style.display = "none"
}

async function handle_professor(obj) {
    const params = new URLSearchParams(window.location.search)
    const dis_div = document.getElementById("adicional");
    const res = await fetch("http://localhost:3000/disciplina/slice", {method: "GET", credentials: "include"})
    const dis_arr = await res.json()

    if (!obj.Disciplinas_ids) {
        obj.Disciplinas_ids = []
    }

    for (const dis of dis_arr) {
        const check = document.createElement("input")
        check.type = "checkbox"

        if (obj.Disciplinas_ids.includes(dis.Id)) {
            check.checked = true
        }

        check.classList.add("disciplina-check")
        check.id = dis.Id
        check.name = dis.Nome
        check.value= dis.Id
        check.id = dis.Nome
        const label = document.createElement("label")
        label.textContent = dis.Nome
        label.for = dis.Nome
        dis_div.appendChild(check)
        dis_div.appendChild(label)
    }
}

function handle_curso(obj) {
    const params = new URLSearchParams(window.location.search)
    document.getElementById("dispo").parentNode.removeChild(document.getElementById("dispo"))

    if (obj.Etapas === null) {
        obj.Etapas = []
    }

    obj.dispo =  undefined

    const ad = document.getElementById("adicional");

    // esses 4 são os 4 que ficam no adicional
    const etapas = document.createElement('div')
    const etapa_counter = document.createElement('div')

    const reload_etapas = () => {
        console.log(obj.Etapas)
        etapas.innerHTML = "" // clear div
        for (let [etidx, et] of obj.Etapas.entries()) {
            // et_ é a cada caixinha que tem várias turmas dentro
            const et_d = document.createElement("div")
            for (const [tidx, t] of et.Turmas.entries()) {
                // child é cada turminha
                const child = document.createElement("div")
                const p = document.createElement("p")
                p.textContent = t.Nome 
                // renomear a turma
                p.onclick = async () => {
                    const str = prompt("Insira o novo nome da turma " + t.Nome)
                    if (!str || str === "" || str.length == 0 ) {
                        return
                    }

                    obj.Etapas[t.Etapa_idx].Turmas[tidx].Nome = str
                    await fetch(
                        `http://localhost:3000/curso/set/${obj.Id}`,
                        {credentials: "include", method: "PUT", body:JSON.stringify(obj)})
                    window.location.reload()
                }

                const bttn = document.createElement("button")
                bttn.textContent = "Ver"
                bttn.onclick = () => {
                    save_item(obj)
                    window.location.replace(`/OtimaHoraWeb/dashboard.html?curso_pai=${obj.Id}&etapa_pai=${t.Etapa_idx}&tipo=turma&$turma=${t.Nome}&id=${t.Id}`);
                }

                const del_bttn = document.createElement("button")
                del_bttn.textContent = "Del"
                del_bttn.onclick = async () => {
                    obj.Etapas[t.Etapa_idx].Turmas = obj.Etapas[t.Etapa_idx].Turmas.splice(tidx, tidx)
                    await fetch(`http://localhost:3000/curso/set/${obj.Id}`, {credentials: "include", method: "PUT", body:JSON.stringify(obj)})
                    window.location.reload()
                }

                child.appendChild(p)
                child.appendChild(bttn)
                child.appendChild(del_bttn)
                child.classList.add("turma")
                et_d.appendChild(child)
                et_d.classList.add("turma-container")
            }
            etapas.classList.add("adicionalDiv")

            const t_count = document.createElement("div")
            t_count.classList.add("turma-counter")
            t_count.textContent = "Turma count: " + et.Turmas.length
            t_count.onclick = async () => {
                const count = parseInt(prompt("Insira a quantidade de turmas para esta etapa"))
                t_count.textContent = "Turma_count: " + count
                et_d.innerHTML = ""
                for (let i = 0; i < count; i++) {
                    const new_t = {Id: 0, Curso_id: obj.Id, Etapa_idx: etidx, Nome:"0", Dispo: Array(5).fill(Array(5).fill(1))}
                    const res = await  fetch(`http://localhost:3000/turma/add`, {
                        credentials: "include",
                        method: "post",
                        body: JSON.stringify(new_t)
                    })
                    const newid = await res.text()
                    new_t.Id = parseInt(newid)
                    new_t['Nome'] = newid
                    et.Turmas.push(new_t)
                }
                et_d.appendChild(t_count)
                reload_etapas()
            }

            const et_bttn = document.createElement("button")
            et_bttn.textContent = "ver"
            et_bttn.onclick = () => {
                save_item(obj)
                window.location.replace(`/OtimaHoraWeb/dashboard.html?curso_pai=${obj.Id}&tipo=etapa&etapa=${etidx}&id=${etidx}`);
            }

            et_d.appendChild(t_count)
            et_d.appendChild(et_bttn)
            etapas.appendChild(et_d)
        }
    }

    etapa_counter.textContent = obj.Etapas.length
    etapa_counter.onclick = () => {
        const str = prompt("novo numero de etapas")
        const new_etapas = []
        for (let i = 0; i < parseInt(str); i++) {
            if (obj.Etapas[i]) {
                new_etapas.push(obj.Etapas[i])
                continue
            }
            new_etapas.push({Curso_id: obj.Id, Idx_in_curso: i, Turmas:[], Curriculo: {}})
        }
        obj.Etapas = new_etapas
        reload_etapas()

        etapa_counter.textContent = str
    }
    etapas.style.display = "flex"
    etapas.style.flexDirection = "row"


    etapa_counter.classList.add("adicionalDiv")

    ad.appendChild(etapa_counter)
    ad.appendChild(etapas)
    ad.style.display = "flex"
    ad.style.flexDirection = "row"
    reload_etapas()
}

async function handle_etapa(obj) {
    const params = new URLSearchParams(window.location.search)
    const response = await fetch(`http://localhost:3000/curso/get/${params.get("curso_pai")}`, {credentials:"include"})
    const curso = await response.json() 

    document.getElementById("dispo").parentNode.removeChild(document.getElementById("dispo"))

    const ad = document.getElementById("adicional");

    const res = await fetch("http://localhost:3000/disciplina/slice", {method: "GET", credentials: "include"})
    const dis_arr = await res.json()

    const materias_sim = document.createElement('div')
    const materias_nao = document.createElement('div')

    let horas_totais = 0
    for (const dis of dis_arr) {
        if (obj.Curriculo[dis.Id] !== undefined) {
            const s = document.createElement('div')
            s.textContent = dis.Nome + ' ' + obj.Curriculo[dis.Id].Horas + " " + obj.Curriculo[dis.Id].Formato
            const btn = document.createElement("button")
            btn.textContent = "remove"
            btn.onclick = async () => {
                obj.Curriculo[dis.Id] = undefined
                curso.Etapas[obj.Idx_in_Curso] = obj

                await fetch(`http://localhost:3000/curso/set/${obj.Curso_id}`, {
                    credentials: "include",
                    method: "PUT",
                    body: JSON.stringify(curso),
                })
                window.location.reload()
            }
            s.appendChild(btn)


            horas_totais += parseInt(obj.Curriculo[dis.Id].Horas)
            materias_sim.appendChild(s)
        } else {
            const s = document.createElement("div")
            s.textContent = dis.Nome
            const btn = document.createElement("button")
            btn.textContent = "add"
            btn.onclick = async () => {
                failed =  true

                let str;
                while (failed == true) {
                    str = prompt("Insira a quantidade de horas e o formato desejado. Exemplo: 10 5+5").split(" ")
                    if (str[1].split("+").map((x) => parseInt(x)).reduce((a, b) => a+b) !== parseInt(str[0])) {
                        alert("Formato não soma para horas")
                        continue
                    }
                    failed = false
                }

                obj.Curriculo[dis.Id] = {Horas: parseInt(str[0]), Formato: str[1]}
                curso.Etapas[obj.Idx_in_Curso].Curriculo[dis.Id] ={Horas: parseInt(str[0]), Formato: str[1]}

                await fetch(`http://localhost:3000/curso/set/${obj.Curso_id}`, {
                    credentials: "include",
                    method: "PUT",
                    body: JSON.stringify(curso),
                })
                window.location.reload()
            }
            s.appendChild(btn)
            materias_nao.appendChild(s)
        }
    }
    materias_sim.appendChild(document.createTextNode(`Horas Totais: ${horas_totais}`))

    materias_sim.style.display = "flex"
    materias_sim.style.flexDirection = "column"
    materias_sim.style.color = "green"
    materias_nao.style.display = "flex"
    materias_nao.style.flexDirection = "column"
    materias_nao.style.color = "red"
    materias_sim.classList.add("adicionalDiv")
    materias_nao.classList.add("adicionalDiv")

    const materias = document.createElement("div")
    const turmas_div = document.createElement("div")

    materias.appendChild(materias_sim)
    materias.appendChild(materias_nao)
    materias.classList.add("materias-container")

    ad.appendChild(materias)
    ad.appendChild(turmas_div)
    ad.style.flexDirection = "column"

    for (const[tidx, t] of obj.Turmas.entries()) {
        // child é cada turminha
        const child = document.createElement("div")
        const p = document.createElement("p")
        p.textContent = t.Nome
        // renomear a turma
        p.onclick = async () => {
            const str = prompt("Insira o novo nome da turma " + t.Nome)
            if (!str || str === "" || str.length == 0 ) {
                return
            }

            curso.Etapas[obj.Idx_in_Curso].Turmas[tidx].Nome = str
            await fetch(
                `http://localhost:3000/curso/set/${obj.Curso_id}`,
                {credentials: "include", method: "PUT", body:JSON.stringify(curso)})
            window.location.reload()
        }

        const bttn = document.createElement("button")
        bttn.textContent = "Ver"
        bttn.onclick = () => {
            save_item(obj)
            window.location.replace(`/OtimaHoraWeb/dashboard.html?curso_pai=${obj.Curso_id}&etapa_pai=${t.Etapa_idx}&tipo=turma&turma=${t.Nome}&id=${t.Id}`);
        }

        const del_bttn = document.createElement("button")
        del_bttn.textContent = "Del"
        del_bttn.onclick = async () => {
            curso.Etapas[obj.Idx_in_Curso].Turmas = curso.Etapas[obj.Idx_in_Curso].Turmas.splice(tidx, tidx)
            await fetch(`http://localhost:3000/curso/set/${obj.Curso_id}`, {credentials: "include", method: "PUT", body:JSON.stringify(curso)})
            window.location.reload()
        }

        child.appendChild(p)
        child.appendChild(bttn)
        child.appendChild(del_bttn)
        child.classList.add("turma")
        turmas_div.appendChild(child)
    }

}

window.get_disp = get_disp
window.set_disp = set_disp
window.save_item = save_item
