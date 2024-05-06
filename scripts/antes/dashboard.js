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

function toggleStatusCol(col_idx) {
    for (let i = 0; i < 5; i ++) {
        toggleStatus(document.getElementById(`${i}-${col_idx}`));
    }
}

function add_professor() {
    const prof_name = document.getElementById("professor-input").value
    document.getElementById("professor-input").value = ""

    const params = new URLSearchParams(window.location.search)

    fetch(`http://localhost:3000/add-professor/${params.get("curso_id")}/${params.get("turma_id")}`, 
        {
            body: JSON.stringify(prof_name),
            method: "POST",
            credentials: "include"
        }
    ).then(window.location.reload())
}

function add_disciplina() {
    const dis_name = document.getElementById("disciplina-input").value
    document.getElementById("disciplina-input").value = ""

    const params = new URLSearchParams(window.location.search)

    fetch(`http://localhost:3000/add-disciplina/${params.get("curso_id")}/${params.get("turma_id")}`, 
        {
            body: JSON.stringify(dis_name),
            method: "POST",
            credentials: "include"
        }
    ).then(window.location.reload())
}

document.addEventListener('DOMContentLoaded', async () => {
   const params = new URLSearchParams(window.location.search)
   const res = await fetch(`http://localhost:3000/session/${params.get("curso_id")}/${params.get("turma_id")}`,
      {
         credentials: "include"
      })

   const turma = await res.json()
   console.log(turma)

   for (const p of turma.Professores) {
      const p_div = document.createElement("div")
      p_div.textContent = p
      document.getElementById("professores").appendChild(p_div)
   }
   for (d of turma.Disciplinas) {
      const d_div = document.createElement("div")
      d_div.textContent = d
      document.getElementById("disciplinas").appendChild(d_div)
   }
})
