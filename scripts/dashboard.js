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
    const obj = {
        segunda: [],
        terca: [],
        quarta: [],
        quinta: [],
        sexta: [],
    }   



    let day_idx = 0
    Object.keys(obj).forEach(day => {
        for (let p_idx = 0; p_idx < PERIODOS; p_idx++) {
            const text = document.getElementById(`${p_idx}-${day_idx}`).textContent
            let value = 1

            if (text === "TALVEZ") {
                value = 0
            }
            if (text === "NAO") {
                value = -1
            }

            obj[day].push(value)
        }
        day_idx += 1
    });

    return obj
}

function set_disp(obj) {
    let day_idx = 0
    Object.keys(obj).forEach(day => {
        for (let p_idx = 0; p_idx < PERIODOS; p_idx++) {
            const value = obj[day][p_idx]
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
        day_idx += 1
    });
}

window.get_disp = get_disp
window.set_disp = set_disp
