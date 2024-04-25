function adcionar_turma(curso) {
   const name = document.getElementById("new_turma").value;
   document.getElementById("new_turma").value = "";
   document.getElementById("turmas").appendChild(new_turma(curso, name));
}

function new_turma(curso, name) {
   let turma = document.createElement("div");

   let c_name = document.createElement("div");
   c_name.textContent = name;
   turma.appendChild(c_name);

   let btn = document.createElement("button");
   btn.onclick = () => {window.location.replace("/turma_dashboard.php?curso=" + curso +"&turma="+name)}
   btn.textContent = "editar";
   turma.appendChild(btn);

   turma.classList.add("item");
   return turma 
}
