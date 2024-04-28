function adcionar_turma() {
   const params = new URLSearchParams(window.location.search)
   const curso = params.get("curso")
   const curso_id = params.get("curso_id")

   const name = document.getElementById("new_turma").value;
   document.getElementById("new_turma").value = "";
   turma = {id: undefined, nome: name}
   fetch(`http://localhost:3000/add-turma/${curso_id}`, {method:"POST", credentials:"include", body:JSON.stringify(turma)})
      .then(response => response.text()) // Get the plain text ID from the response
      .then(turma_id => {
         document.getElementById("turmas").appendChild(new_turma(curso, curso_id, name, turma_id));
      });
}

function new_turma(curso, curso_id, name, turma_id) {
   let turma = document.createElement("div");

   let c_name = document.createElement("div");
   c_name.textContent = name;
   turma.appendChild(c_name);

   let btn = document.createElement("button");
  
   btn.onclick = () => {window.location.replace(`/turma_dashboard.php?curso=${curso}&curso_id=${curso_id}turma=${name}&turma_id=${turma_id}`)}
   btn.textContent = "editar";
   turma.appendChild(btn);

   turma.classList.add("item");
   return turma 
}
