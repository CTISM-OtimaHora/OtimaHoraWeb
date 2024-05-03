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


function add_professor() {
    const params = new URLSearchParams(window.location.search)
    const name = document.getElementById("professor-input").value
    fetch(
        `http://localhost:3000/add-professor/${params.get("curso_id")}/${params.get("turma_id")}`
        , {method:"POST", credentials:'include', body:JSON.stringify(name)}).then()
   document.getElementById("professor-input").value = ""
}

function add_disciplina() {
   const params = new URLSearchParams(window.location.search)
   const name = document.getElementById("disciplina-input").value
   fetch(
      `http://localhost:3000/add-disciplina/${params.get("curso_id")}/${params.get("turma_id")}`
      , {method:"POST", credentials:'include', body:JSON.stringify(name)}).then()
   document.getElementById("professor-input").value = ""
}

function toggleStatusCol(col_idx) {
    for (let i = 0; i < 5; i ++) {
        toggleStatus(document.getElementById(`${i}-${col_idx}`));
    }
}
