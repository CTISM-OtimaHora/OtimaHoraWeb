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
    ).then()
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
    ).then()
}
