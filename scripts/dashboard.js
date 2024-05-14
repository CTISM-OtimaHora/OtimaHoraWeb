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


function save_disp() {
    const params = new URLSearchParams(window.location.search)
    const obj = get_disp()

     fetch(`http://localhost:3000/set-disp/${params.get("tipo")}/${params.get("id")}`, 
        {
            credentials: "include",
            method: "POST",
            body: JSON.stringify(obj)
        }).then(alert("saved"))
}

document.addEventListener('DOMContentLoaded', async function() {
    const params = new URLSearchParams(window.location.search)
    let disp;

    if (params.get("tipo") == "contrato") {
        const res = await fetch(`http://localhost:3000/session/get/contrato/${params.get("id")}`, 
            {
                credentials: "include",
                method: "GET",
            })
        const cont = await res.json()
        disp = cont.Dispo

        const part = document.getElementById("participantes")
        part.style.display = "flex"
        part.style.flexDirection = "column"

        for (const p of cont.Participantes) {
            console.log(p)
            const child = document.createElement("div")
            child.style.display = "flex"
            child.textContent = `${p.Tipo} - ${p.Nome}`

            const bttn = document.createElement("button")
            bttn.textContent = "Ver"
            bttn.onclick = () => {
                window.location.replace(`/dashboard.html?tipo=${p.Tipo}&${p.Tipo}=${p.Nome}&id=${p.Id}`);
            }
            child.appendChild(bttn)
            part.appendChild(child)
        }
        document.getElementById("save").style.display = "none"
    } else {
        const res = await fetch(`http://localhost:3000/get-disp/${params.get("tipo")}/${params.get("id")}`, 
            {
                credentials: "include",
                method: "GET",
            })

        disp = await res.json()
    }

    set_disp(disp)
})

window.get_disp = get_disp
window.set_disp = set_disp

window.save_disp = save_disp
