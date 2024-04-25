function adcionar_curso() {
   const name = document.getElementById("new_curso").value;
   document.getElementById("new_curso").value = "";
   document.getElementById("cursos").appendChild(new_curso(name));
}

function new_curso(name) {
   let curso = document.createElement("div");

   let c_name = document.createElement("div");
   c_name.textContent = name;
   curso.appendChild(c_name);

   let btn = document.createElement("button");
   btn.onclick = () => {window.location.replace("/turmas.php?curso=" + name)}
   btn.textContent = "turmas";
   curso.appendChild(btn);

   curso.classList.add("item");
   curso.classList.add("curso");
   return curso
}
